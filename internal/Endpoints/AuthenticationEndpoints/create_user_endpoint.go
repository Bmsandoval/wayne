package AuthenticationEndpoints

import (
	"context"
	"github.com/bmsandoval/wayne/internal/Services"
	"github.com/bmsandoval/wayne/pkg/AppContext"
	"github.com/bmsandoval/wayne/protos/authenticator"
	"github.com/go-kit/kit/endpoint"
	"log"
)

func MakeCreateUserEndpoint(appCtx AppContext.Context, services Services.Bundle) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		in := request.(*authenticator.CreateUserRequest)

		uid, err := services.UserSvc.Create(in.Username, in.Password)
		if err != nil {
			log.Println(err.Error())
			return &authenticator.CreateUserResponse{
				Success: false,
			}, nil
		}
		_=uid

		return &authenticator.CreateUserResponse{
			Success: true,
		}, nil
	}
}
