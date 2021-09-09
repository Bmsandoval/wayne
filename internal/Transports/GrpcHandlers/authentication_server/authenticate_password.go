package authentication_server

import (
	"context"
	"github.com/bmsandoval/wayne/internal/Endpoints/AuthenticationEndpoints"
	"github.com/bmsandoval/wayne/protos/authenticator"
)

func (s *Server) AuthenticatePassword(ctx context.Context, in *authenticator.AuthenticatePasswordRequest) (*authenticator.AuthenticatePasswordResponse, error) {
	endpoint := AuthenticationEndpoints.MakeAuthenticatePasswordEndpoint(s.Context, s.Bundle)

	response, err := endpoint(ctx, in)

	if err != nil {
		return nil, err
	}

	if response == nil {
		return nil, nil
	}

	return response.(*authenticator.AuthenticatePasswordResponse), nil
}
