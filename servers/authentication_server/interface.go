package authentication_server

import (
	"github.com/bmsandoval/wayne/servers"
	"google.golang.org/grpc"
)

// server is used to implement helloworld.GreeterServer.
type Server struct {
	servers.ServerContext
}

//type Registerable struct {}
func init() {
	servers.Include(func(s *grpc.Server, ctx servers.ServerContext){
		RegisterAuthenticatorServer(s, &Server{
			ServerContext: ctx,
		})
	})
}
