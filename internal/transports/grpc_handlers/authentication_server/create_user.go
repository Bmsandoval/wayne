package authentication_server

import (
	"context"
	"github.com/bmsandoval/wayne/protos/authenticator"
)

func (s *Server) CreateUser(ctx context.Context, in *authenticator.CreateUserRequest) (*authenticator.CreateUserResponse, error) {
	//endpoint := authentication_endpoints.MakeVehicleReadEndpoint(s.AppCtx, s.Bundle)
	//response, err := endpoint(ctx, in)
	////response, err := endpoint(s.AppCtx.GoContext, in)
	//if err != nil {
	//	return nil, err
	//}
	//if response == nil {
	//	return nil, nil
	//}
	//return response.(*authenticator.CreateUserResponse), nil
	return nil, nil
}
