package authentication_server

import (
	"context"
	"log"
)

func (s *Server) AuthenticatePassword(ctx context.Context, in *AuthenticatePasswordRequest) (*AuthenticatePasswordResponse, error) {
	sub, err := s.Bundle.UserSvc.ValidatePassword(in.Username, in.Password)
	log.Println(sub)
	if err != nil {
		return &AuthenticatePasswordResponse{
			Success: false,
		}, nil
	}

	return &AuthenticatePasswordResponse{
		Success: true,
	}, nil
}
