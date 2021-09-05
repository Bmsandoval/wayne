package grpc_handlers

import (
	"github.com/bmsandoval/wayne/internal/service"
	"github.com/bmsandoval/wayne/internal/utilities/appcontext"
	"google.golang.org/grpc"
)


type RegisterableGrpcHandler func(*grpc.Server, appcontext.Context, service.Bundle)
var RegisterableGrpcHandlers []RegisterableGrpcHandler
func RegisterGrpcHandler(registerableGrpcHandler RegisterableGrpcHandler) {
	RegisterableGrpcHandlers = append(RegisterableGrpcHandlers, registerableGrpcHandler)
}


func ConfigureGrpcHandlers(server *grpc.Server, appCtx appcontext.Context, bundle service.Bundle) {
	for _, registerableHandler := range RegisterableGrpcHandlers {
		registerableHandler(server, appCtx, bundle)
	}
}