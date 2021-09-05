package Entry

import (
	"fmt"
	"github.com/bmsandoval/wayne/configs"
	"github.com/bmsandoval/wayne/internal/Db"
	"github.com/bmsandoval/wayne/internal/Services"
	"github.com/bmsandoval/wayne/internal/Transports/GrpcHandlers"
	"github.com/bmsandoval/wayne/internal/Transports/HttpHandlers"
	"github.com/bmsandoval/wayne/pkg/AppContext"
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
	connection, err := Db.Start(*config)
	if err != nil {
		panic(err) }
	defer func() {
		if err := Db.Stop(); err != nil {
			panic(err) }
	}()

	// Build Context
	ctx := AppContext.Context{
		Config: *config,
		DB: *connection,
		// Redis
	}

	// Bundle Services
	serviceBundle, err := Services.NewBundle(ctx)
	if err != nil {
		panic(err) }

	// Bundle Servers
	grpcS := grpc.NewServer()
	GrpcHandlers.ConfigureGrpcHandlers(grpcS, ctx, *serviceBundle)

	router := mux.NewRouter()
	httpS := &http.Server{
		Handler: router,
	}
	HttpHandlers.ConfigureHttpHandlers(router, ctx, *serviceBundle)

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
