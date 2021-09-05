package authentication_server

import (
	"context"
	"github.com/bmsandoval/wayne/protos/authenticator"
)

func (s *Server) AuthenticatePassword(ctx context.Context, in *authenticator.AuthenticatePasswordRequest) (*authenticator.AuthenticatePasswordResponse, error) {
	//user, err := s.Bundle.UserSvc.ValidatePassword(in.Username, in.Password, models.User{})
	//if err != nil {
	//	log.Println(err.Error())
	//	return &authenticator.AuthenticatePasswordResponse{
	//		Success: false,
	//	}, nil
	//}
	//
	//if user == nil {
	//	log.Println("here")
	//}
	//
	//return &authenticator.AuthenticatePasswordResponse{
	//	Success: true,
	//}, nil
	return nil, nil
}
