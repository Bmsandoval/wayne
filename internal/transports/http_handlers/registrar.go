package http_handlers

import (
	"github.com/bmsandoval/wayne/internal/service"
	"github.com/bmsandoval/wayne/internal/utilities/appcontext"
	"github.com/gorilla/mux"
)


type RegisterableHttpHandler func(*mux.Router, appcontext.Context, service.Bundle)
var RegisterableHttpHandlers []RegisterableHttpHandler
func RegisterHttpHandler(registerableHttpHandler RegisterableHttpHandler) {
	RegisterableHttpHandlers = append(RegisterableHttpHandlers, registerableHttpHandler)
}


func ConfigureHttpHandlers(server *mux.Router, appCtx appcontext.Context, bundle service.Bundle) {
	for _, registerableHandler := range RegisterableHttpHandlers {
		registerableHandler(server, appCtx, bundle)
	}
}