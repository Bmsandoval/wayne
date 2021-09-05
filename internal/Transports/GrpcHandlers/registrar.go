package GrpcHandlers

import (
	"github.com/bmsandoval/wayne/internal/Services"
	"github.com/bmsandoval/wayne/pkg/AppContext"
	"google.golang.org/grpc"
)


type RegisterableGrpcHandler func(*grpc.Server, AppContext.Context, Services.Bundle)
var RegisterableGrpcHandlers []RegisterableGrpcHandler
func RegisterGrpcHandler(registerableGrpcHandler RegisterableGrpcHandler) {
	RegisterableGrpcHandlers = append(RegisterableGrpcHandlers, registerableGrpcHandler)
}


func ConfigureGrpcHandlers(server *grpc.Server, appCtx AppContext.Context, bundle Services.Bundle) {
	for _, registerableHandler := range RegisterableGrpcHandlers {
		registerableHandler(server, appCtx, bundle)
	}
}