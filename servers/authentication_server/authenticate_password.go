package authentication_server

import (
	"context"
	"github.com/bmsandoval/wayne/db/models"
	"log"
)

func (s *Server) AuthenticatePassword(ctx context.Context, in *AuthenticatePasswordRequest) (*AuthenticatePasswordResponse, error) {
	user, err := s.Bundle.UserSvc.ValidatePassword(in.Username, in.Password, models.User{})
	if err != nil {
		log.Println(err.Error())
		return &AuthenticatePasswordResponse{
			Success: false,
		}, nil
	}

	if user == nil {
		log.Println("here")
	}

	return &AuthenticatePasswordResponse{
		Success: true,
	}, nil
}
