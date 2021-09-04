package authentication_server

import (
	"context"
	"log"
)

func (s *Server) UsernameAvailable(ctx context.Context, in *UsernameAvailableRequest) (*UsernameAvailableResponse, error) {
	available, err := s.Bundle.UserSvc.UsernameAvailable(in.Username)
	if err != nil {
		log.Println(err.Error())
		return &UsernameAvailableResponse{
			Available: false,
		}, nil
	}

	return &UsernameAvailableResponse{
		Available: available,
	}, nil
}
