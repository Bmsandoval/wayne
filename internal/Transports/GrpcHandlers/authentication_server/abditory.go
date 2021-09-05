package authentication_server

import (
	"github.com/bmsandoval/wayne/internal/Services"
	"github.com/bmsandoval/wayne/internal/Transports/GrpcHandlers"
	"github.com/bmsandoval/wayne/pkg/AppContext"
	"github.com/bmsandoval/wayne/protos/authenticator"
	"google.golang.org/grpc"
)

func init() {
	GrpcHandlers.RegisterGrpcHandler(
		func(s *grpc.Server, ctx AppContext.Context, bundle Services.Bundle){
			authenticator.RegisterAuthenticatorServer(s,
				&Server{
				ctx,bundle,
				})
	})
}

type Server struct {
	AppContext.Context
	Services.Bundle
}

