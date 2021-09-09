package auth_http_handlers

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/bmsandoval/wayne/internal/Endpoints/AuthenticationEndpoints"
	"github.com/bmsandoval/wayne/internal/Services"
	"github.com/bmsandoval/wayne/internal/Transports/HttpHandlers"
	"github.com/bmsandoval/wayne/pkg/AppContext"
	"github.com/bmsandoval/wayne/protos/authenticator"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

func init() {
	HttpHandlers.RegisterHttpHandler(MakeAuthenticatePasswordHttpHandler())
}

func MakeAuthenticatePasswordHttpHandler() HttpHandlers.RegisterableHttpHandler {
	return func(router *mux.Router, appCtx AppContext.Context, services Services.Bundle) {
		api := router.PathPrefix("/auth").Subrouter()

		endpoint := AuthenticationEndpoints.MakeAuthenticatePasswordEndpoint(appCtx, services)
		decoder, _ := MakeAuthenticatePasswordRequestDecoder(appCtx)
		encoder, _ := MakeAuthenticatePasswordResponseEncoder(appCtx)

		api.Methods("GET").Path("/user").Handler(kithttp.NewServer(
			endpoint,
			decoder,
			encoder,
		))
	}
}

func MakeAuthenticatePasswordRequestDecoder(appCtx AppContext.Context) (kithttp.DecodeRequestFunc, error) {
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		var Req authenticator.AuthenticatePasswordRequest

		// r.Body is an io.ReadCloser, need to decode
		decode := json.NewDecoder(r.Body)
		err := decode.Decode(&Req)
		if err != nil {
			return nil, err
		}

		return &Req, nil
	}, nil
}

func MakeAuthenticatePasswordResponseEncoder(appCtx AppContext.Context) (kithttp.EncodeResponseFunc, error) {
	return func(ctx context.Context, httpResponse http.ResponseWriter, endpointResponse interface{}) error {
		// DECODE RESPONSE
		resp, ok := endpointResponse.(*authenticator.AuthenticatePasswordResponse)
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
