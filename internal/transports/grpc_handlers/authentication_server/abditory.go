package authentication_server

import (
	"github.com/bmsandoval/wayne/internal/service"
	"github.com/bmsandoval/wayne/internal/transports/grpc_handlers"
	"github.com/bmsandoval/wayne/internal/utilities/appcontext"
	"github.com/bmsandoval/wayne/protos/authenticator"
	"google.golang.org/grpc"
)

func init() {
	grpc_handlers.RegisterGrpcHandler(
		func(s *grpc.Server, ctx appcontext.Context, bundle service.Bundle){
			authenticator.RegisterAuthenticatorServer(s,
				&Server{
				ctx,bundle,
				})
	})
}

type Server struct {
	appcontext.Context
	service.Bundle
}

