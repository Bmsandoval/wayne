Table of contents
=================

* [Table of contents](#table-of-contents)
* [Overview](#overview)
* [Requirements](#requirements)
* [Setup](#setup)
  * [Dev Environment](#setup-dev-environment)
    * [1) Pull Down Repo](#1-pull-down-the-repo)
    * [2) Install Golang](#2-install-golang)
    * [3) Configure Golang](#3-configure-golang)
    * [3) Configure Golang](#3-configure-golang)
    * [4) Download Dependencies](#4-download-dependencies)
    * [5) Install Protoc](#5-install-protoc)
    * [6) Install DockerCE](#6-install-dockerce)
    * [7) Source Bash Profile](#7-source-the-bash-profile)
  
# Overview

Wayne (Way-In) is an open-source, self-hosted, auth solution written as a GoLang-based micro-service. Local development available via Docker. Wayne is the toughest guy in LetterKenny, let him be your bouncer.

---

# Setup

## Setup Dev Environment

> ### 1) Pull down the repo
>
> ```
> git clone git@github.com:bmsandoval/wayne.git
> ```

> ### 2) Install Golang 
> * https://golang.org/dl/
> * Recommended - go version go1.13.1 darwin/amd64)

> ### 3) Configure Golang
> * add the following to your bash profile. these may very.
>   * IMPORTANT : since using modules, this project must NOT reside in your GOPATH listed below
> ```
> export PATH=$PATH:/usr/local/go/bin:$GOBIN
> export GOROOT=/usr/local/opt/go/libexec
> export PATH=$PATH:$GOROOT/bin
> export GOPATH=~/projects/go
> export PATH=$PATH:$GOPATH/bin
> ```

> ### 4) Download Dependencies
> ```
> go mod download && go mod vendor
> ```

> ### 5) Install Protoc
> ```
> apt install -y protobuf-compiler
> ```
> and set it up with your global golang
> ```
> go get -u github.com/golang/protobuf/protoc-gen-go
> ```

> ### 6) Install DockerCE 
> * Recommended - Docker version 20.10.7

> ### 7) Source the bash profile
> ```
> . .profile
> ```
