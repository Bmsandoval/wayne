package auth_http_handlers

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/bmsandoval/wayne/internal/endpoints/authentication_endpoints"
	"github.com/bmsandoval/wayne/internal/service"
	"github.com/bmsandoval/wayne/internal/transports/http_handlers"
	"github.com/bmsandoval/wayne/internal/utilities/appcontext"
	"github.com/bmsandoval/wayne/protos/authenticator"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

func init() {
	http_handlers.RegisterHttpHandler(MakeUsernameAvailableHttpHandler())
}

func MakeUsernameAvailableHttpHandler() http_handlers.RegisterableHttpHandler {
	return func(router *mux.Router, appCtx appcontext.Context, services service.Bundle) {
		api := router.PathPrefix("/auth").Subrouter()

		endpoint := authentication_endpoints.MakeUsernameAvailableEndpoint(appCtx, services)
		decoder, _ := MakeUsernameAvailableRequestDecoder(appCtx)
		encoder, _ := MakeUsernameAvailableResponseEncoder(appCtx)

		api.Methods("GET").Path("/username/available").Handler(kithttp.NewServer(
			endpoint,
			decoder,
			encoder,
		))
	}
}

func MakeUsernameAvailableRequestDecoder(appCtx appcontext.Context) (kithttp.DecodeRequestFunc, error) {
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		var Req authenticator.UsernameAvailableRequest

		// r.Body is an io.ReadCloser, need to decode
		decode := json.NewDecoder(r.Body)
		err := decode.Decode(&Req)
		if err != nil {
			return nil, err
		}

		return &Req, nil
	}, nil
}

func MakeUsernameAvailableResponseEncoder(appCtx appcontext.Context) (kithttp.EncodeResponseFunc, error) {
	return func(ctx context.Context, httpResponse http.ResponseWriter, endpointResponse interface{}) error {
		// DECODE RESPONSE
		resp, ok := endpointResponse.(*authenticator.UsernameAvailableResponse)
		if ! ok {
			return errors.New("response was the wrong type")
		}
		// MARSHAL RESPONSE
		e, err := json.Marshal(resp)
		if err != nil {
			return err
		}
		// WRITE RESPONSE
		_, err = httpResponse.Write(e)
		return err
	}, nil
}
