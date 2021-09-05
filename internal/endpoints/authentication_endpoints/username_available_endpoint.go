package authentication_endpoints

import (
	"context"
	"github.com/bmsandoval/wayne/internal/service"
	"github.com/bmsandoval/wayne/internal/utilities/appcontext"
	"github.com/bmsandoval/wayne/protos/authenticator"
	"github.com/go-kit/kit/endpoint"
	"log"
)

func MakeUsernameAvailableEndpoint(appCtx appcontext.Context, services service.Bundle) endpoint.Endpoint {
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
