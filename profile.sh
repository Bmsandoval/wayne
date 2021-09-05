#!/usr/bin/env bash

# Just gets the top level directory of this project. Useful for scripting within the project via relative file paths
WAYNE_CODE_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

_wayne_base_options=\
"protoc\t:\tGenerate grpc files from the proto files
mock\t:\tMock all services (req: gomock)
test\t:\tRun all mock tests
start\t:\tLaunch wayne's docker container(s). Optionally specify specific one
stop\t:\tStop wayne's docker container(s)
remake\t:\tForce rebuild. probably don't need to do this, it's built into the start and reset commands
reset\t:\tPurges and rebuilds wayne's docker container(s) [[WARNING: RESETS DATABASE]]
purge\t:\tStop containers and purge all remnants [[WARNING: EVEN WORSE THAN RESET]]
help\t:\tShow this menu
edit\t:\tHelper for opening this file in vim
src\t:\tRe-Source this file"

_wayne_purge_options=\
"api\t:\tThis project's backend's docker container (db unaffected)
project\t:\tAll docker containers for this project
everywhere\t:\tAll non-database docker containers, regardless of project
everything\t:\tAll docker containers, regardless of project
killme\t:\tPurges base docker image. You probably don't want to do this
help\t:\tDisplays this menu"

_wayne_start_stop_options=\
"api\t:\tThis project's backend's docker container (db unaffected)
db\t:\tThis project's database's docker container (db unaffected)
project\t:\tAll docker containers for this project
help\t:\tDisplays this menu"

alias we="wayne"
wayne () {
  local _baseOption="${1}"
  shift
  local _subOption="${1}"
  case "${_baseOption}" in
    'help'|'')
      echo -e "Usage: $ ${FUNCNAME[0]} [option]
Options:
${_wayne_base_options}"
    ;;
    'src') . "${WAYNE_CODE_DIR}/profile.sh";;
    'start')
      case "${_subOption}" in
        'api') docker compose -f "${WAYNE_CODE_DIR}/docker-compose.yml" up --build wayne-api ;;
        'db') docker compose -f "${WAYNE_CODE_DIR}/docker-compose.yml" up -d wayne-db ;;
        'project') ${FUNCNAME[0]} start db && ${FUNCNAME[0]} start api ;;
        ''|*) echo -e "Usage: $ ${FUNCNAME[0]} ${_baseOption} [option]
Options:
${_wayne_start_stop_options}"
      esac
      echo "STARTING wayne services"
    ;;
    'stop')
      echo "STOPPING wayne services"
      case "${_subOption}" in
        'api') docker compose -f "${WAYNE_CODE_DIR}/docker-compose.yml" down wayne-api ;;
        'db') docker compose -f "${WAYNE_CODE_DIR}/docker-compose.yml" down wayne-db ;;
        'project') docker compose -f "${WAYNE_CODE_DIR}/docker-compose.yml" down; ${FUNCNAME[0]} sync -cease ;;
        ''|*) echo -e "Usage: $ ${FUNCNAME[0]} ${_baseOption} [option]
Options:
${_wayne_start_stop_options}"
      esac
    ;;
    'rebuild')
      docker compose -f "${WAYNE_CODE_DIR}/App/docker compose.yml" build
    ;;
    'purge')
      case "${_subOption}" in
        'api') ${FUNCNAME[0]} stop api; docker system prune --all --force --filter=label=base=true --filter=label=notdb=true --filter=label=wayne=true ;;
        'project') docker stop $(docker ps -q); docker system prune --all --force --filter=label=base=true --filter=label=wayne=true ;;
        'everywhere') docker stop $(docker ps -q); docker system prune --all --force --filter=label=base=true --filter=label=notdb=true ;;
        'everything') docker stop $(docker ps -q); docker system prune --all --force --filter=label=base=true ;;
        'killme') docker stop $(docker ps -q); docker system prune --all --force ;;
        ''|*) echo -e "Usage: $ ${FUNCNAME[0]} ${_baseOption} [option]
Options:
${_wayne_purge_options}"
      esac
    ;;
    'reset')
      case "${_subOption}" in
        'api') ${FUNCNAME[0]} purge api; ${FUNCNAME[0]} rebuild ;;
        'project') ${FUNCNAME[0]} purge project; ${FUNCNAME[0]} rebuild ;;
        ''|*) echo "'api' or 'project'"
      esac
    ;;
    'protoc')
      wayneProtoc
    ;;
    'mock')
      wayneMockServices
    ;;
    'test')
      wayneTestServers
    ;;
    *)
      echo -e "ERROR: invalid option. Try..\n$ ${FUNCNAME} help"
    ;;
  esac
}


wayneTestServers () {
 go test $(go list $WAYNE_CODE_DIR/...)
}


wayneProtoc () {
  packageName="protoc"
  package-installed "${packageName}"

  # Note that in bash, non-zero exit codes are error codes. returning 0 means success
  if [[ "$?" == "0" ]]; then
    # If installed, run protoc
    PROTO_FOLDER="protos"
    SERVER_DIR="${WAYNE_CODE_DIR}/${PROTO_FOLDER}"
    # need relative path. cd in subshell to have fine return a path relative to the proto folder
    SERVERS=$(cd "${SERVER_DIR}" && find . -maxdepth 1 -mindepth 1 -type d)
    for SERVER in ${SERVERS}; do
      # for each server found, run proto
      protoc --go_out=plugins=grpc:. "${PROTO_FOLDER}/${SERVER}"/*.proto
    done
  else
    # If protobuf missing, tell them to install it
    echo "missing required package 'protobuf'. Please run the following commands and try again:"
    echo "install protobuf, and then run..."
    echo "$ go get -u github.com/golang/protobuf/protoc-gen-go"
  fi
}


# Generate mock files for all service, putting the results in the proper file. renames some stuff for consistency.
# If you update any service, recommend running this function to update the service for the tests.
wayneMockServices () {
  packageName="mockgen"
  which "${packageName}"

  # Note that in bash, non-zero exit codes are error codes. returning 0 means success
  if [[ "$?" == "0" ]]; then
    MOCK_FOLDER="service"
    SERVICE_DIR="${WAYNE_CODE_DIR}/internal/${MOCK_FOLDER}"
    SERVICES=$(find "${SERVICE_DIR}" -maxdepth 1 -mindepth 1 -type d)
    for SERVICE_PATH in ${SERVICES}
    do
      if [[ -f ${SERVICE_PATH}/abditory.go ]]; then
        FOLDER_NAME="${SERVICE_PATH##*/}"
        mockgen \
          -source=${SERVICE_PATH}/abditory.go \
          -destination=mocks/${MOCK_FOLDER}_mocks/${FOLDER_NAME}_mock.go \
          -package=${MOCK_FOLDER}_mocks \
          -mock_names Service=Mock_${FOLDER_NAME}
        fi
    done
  else
    # If mockgen missing, tell them to install it
    echo "missing required package 'mockgen'. Please run the following commands and try again:"
    echo "install protobuf, and then run..."
    echo "$ go get -u github.com/golang/protobuf/protoc-gen-go"
  fi
}


# Check if a command exists in the environment
# Returns 0 if command found
package-installed () {
	result=$(compgen -A function -abck | grep "^$1$")
  # Note that in bash, non-zero exit codes are error codes. returning 0 means success
	if [[ "${result}" == "$1" ]]; then
		# package installed
		return 0
	else
		# package not installed
		return 1
	fi
}


_fzf_complete_wayne () {
  #####################
  ## HANDLE BASE OPTIONS
  if [[ "${COMP_CWORD}" == "1" ]]; then
    if [[ "${COMP_WORDS[COMP_CWORD]}" == "**" ]]; then
      which fzf >/dev/null && _fzf_complete --with-nth=1 --delimiter='\t' --preview='echo -e {3}' --preview-window=up:sharp:wrap:40% --prompt="wayne> " -- "$@" < <(
        echo -e "${_wayne_base_options}")
    else
      COMPREPLY=($(compgen -W "$(echo -e "${_wayne_base_options}" | perl -ne 'print "$1 " if /([^\t]+)\t/')" "${COMP_WORDS[COMP_CWORD]}"))
    fi
  elif [[ "${COMP_WORDS[COMP_CWORD]}" != "**" ]]; then
    #####################
    ## HANDLE NORMAL OPTIONS
    case "${COMP_WORDS[COMP_CWORD-1]}" in
    "start"|"stop") COMPREPLY=("" $(compgen -W "$(echo -e "${_wayne_start_stop_options}" | perl -ne 'print "$1 " if /([^\t]+)\t/')" "${COMP_WORDS[COMP_CWORD]}")) ;;
    "purge") COMPREPLY=("" $(compgen -W "$(echo -e "${_wayne_purge_options}" | perl -ne 'print "$1 " if /([^\t]+)\t/')" "${COMP_WORDS[COMP_CWORD]}")) ;;
    "reset") COMPREPLY=($(compgen -W "api project" "${COMP_WORDS[COMP_CWORD]}")) ;;
    esac
  else
  #####################
  ## HANDLE FZF OPTIONS
    case "${COMP_WORDS[COMP_CWORD-1]}" in
    "purge")
      which fzf >/dev/null && _fzf_complete --with-nth=1 --delimiter='\t' --preview='echo -e {3}' --preview-window=up:sharp:wrap:40% --prompt="wayne> " -- "$@" < <(
        echo -e "${_wayne_purge_options}")
    ;;
    "start"|"stop")
      which fzf >/dev/null && _fzf_complete --with-nth=1 --delimiter='\t' --preview='echo -e {3}' --preview-window=up:sharp:wrap:40% --prompt="wayne> " -- "$@" < <(
        echo -e "${_wayne_start_stop_options}")
    ;;
    esac
  fi
}


complete -F _fzf_complete_wayne -o default -o bashdefault wayne
complete -F _fzf_complete_wayne -o default -o bashdefault we