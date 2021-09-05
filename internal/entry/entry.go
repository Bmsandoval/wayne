package entry

import (
	"fmt"
	"github.com/bmsandoval/wayne/configs"
	"github.com/bmsandoval/wayne/internal/db"
	"github.com/bmsandoval/wayne/internal/service"
	"github.com/bmsandoval/wayne/internal/transports/grpc_handlers"
	"github.com/bmsandoval/wayne/internal/transports/http_handlers"
	"github.com/bmsandoval/wayne/internal/utilities/appcontext"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	//gocache "github.com/pmylund/go-cache"
	"github.com/soheilhy/cmux"
)

func Entry() {
	// Get Configs
	config, err := configs.Configure()
	if err != nil {
		panic(err) }

	// Setup Database
	connection, err := db.Start(*config)
	if err != nil {
		panic(err) }
	defer func() {
		if err := db.Stop(); err != nil {
			panic(err) }
	}()

	// Build Context
	ctx := appcontext.Context{
		Config: *config,
		DB: *connection,
		// Redis
	}

	// Bundle Services
	serviceBundle, err := service.NewBundle(ctx)
	if err != nil {
		panic(err) }

	// Bundle Servers
	grpcS := grpc.NewServer()
	grpc_handlers.ConfigureGrpcHandlers(grpcS, ctx, *serviceBundle)

	router := mux.NewRouter()
	httpS := &http.Server{
		Handler: router,
	}
	http_handlers.ConfigureHttpHandlers(router, ctx, *serviceBundle)

	// Create the main listener.
	l, err := net.Listen("tcp", fmt.Sprintf(":%s", config.SrvPort))
	if err != nil {
		log.Fatal(err)
	}

	// Create a cmux.
	m := cmux.New(l)

	// Match connections in order:
	// First grpc, then HTTP
	grpcL := m.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
	httpL := m.Match(cmux.HTTP1Fast())

	// Use the muxed listeners for your servers.
	go grpcS.Serve(grpcL)
	go httpS.Serve(httpL)

	// Start serving!
	log.Println("Starting Server...")
	m.Serve()
}
