package authentication_server

import (
	"context"
)

func (s *Server) CreateUser(ctx context.Context, in *CreateUserRequest) (*CreateUserResponse, error) {
	_, err := s.Bundle.UserSvc.Create(in.Username, in.Password)
	if err != nil {
		return &CreateUserResponse{
			Success: true,
		}, nil
	}

	return &CreateUserResponse{
		Success: true,
	}, nil
}
