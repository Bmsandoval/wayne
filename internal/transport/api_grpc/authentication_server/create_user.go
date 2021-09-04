package authentication_server

import (
	"context"
	"log"
)

func (s *Server) CreateUser(ctx context.Context, in *CreateUserRequest) (*CreateUserResponse, error) {
	_, err := s.Bundle.UserSvc.Create(in.Username, in.Password)
	if err != nil {
		log.Println(err.Error())
		return &CreateUserResponse{
			Success: false,
		}, nil
	}

	return &CreateUserResponse{
		Success: true,
	}, nil
}
