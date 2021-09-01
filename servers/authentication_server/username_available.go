package authentication_server

import (
	"context"
)

func (s *Server) UsernameAvailable(ctx context.Context, in *UsernameAvailableRequest) (*UsernameAvailableResponse, error) {
	available, err := s.Bundle.UserSvc.UsernameAvailable(in.Username)
	if err != nil {
		return &UsernameAvailableResponse{
			Available: false,
		}, nil
	}

	return &UsernameAvailableResponse{
		Available: available,
	}, nil
}
