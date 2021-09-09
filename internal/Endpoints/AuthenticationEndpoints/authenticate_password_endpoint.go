package AuthenticationEndpoints

import (
	"context"
	"github.com/bmsandoval/wayne/internal/Db/models"
	"github.com/bmsandoval/wayne/internal/Services"
	"github.com/bmsandoval/wayne/pkg/AppContext"
	"github.com/bmsandoval/wayne/protos/authenticator"
	"github.com/go-kit/kit/endpoint"
	"log"
)

func MakeAuthenticatePasswordEndpoint(appCtx AppContext.Context, services Services.Bundle) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		in := request.(*authenticator.AuthenticatePasswordRequest)

		_, err := services.UserSvc.ValidatePassword(in.Username, in.Password, models.User{})
		if err != nil {
			log.Println(err.Error())
			return &authenticator.AuthenticatePasswordResponse{
				Success: false,
			}, nil
		}

		return &authenticator.AuthenticatePasswordResponse{
			Success: true,
		}, nil
	}
}
