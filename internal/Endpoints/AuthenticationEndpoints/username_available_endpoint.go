package AuthenticationEndpoints

import (
	"context"
	"github.com/bmsandoval/wayne/internal/Services"
	"github.com/bmsandoval/wayne/pkg/AppContext"
	"github.com/bmsandoval/wayne/protos/authenticator"
	"github.com/go-kit/kit/endpoint"
	"log"
)

func MakeUsernameAvailableEndpoint(appCtx AppContext.Context, services Services.Bundle) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		in := request.(*authenticator.UsernameAvailableRequest)

		available, err := services.UserSvc.UsernameAvailable(in.Username)
		if err != nil {
			log.Println(err.Error())
			return &authenticator.UsernameAvailableResponse{
				Available: false,
			}, nil
		}

		return &authenticator.UsernameAvailableResponse{
			Available: available,
		}, nil
	}
}
