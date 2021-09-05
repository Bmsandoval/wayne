package HttpHandlers

import (
	"github.com/bmsandoval/wayne/internal/Services"
	"github.com/bmsandoval/wayne/pkg/AppContext"
	"github.com/gorilla/mux"
)


type RegisterableHttpHandler func(*mux.Router, AppContext.Context, Services.Bundle)
var RegisterableHttpHandlers []RegisterableHttpHandler
func RegisterHttpHandler(registerableHttpHandler RegisterableHttpHandler) {
	RegisterableHttpHandlers = append(RegisterableHttpHandlers, registerableHttpHandler)
}


func ConfigureHttpHandlers(server *mux.Router, appCtx AppContext.Context, bundle Services.Bundle) {
	for _, registerableHandler := range RegisterableHttpHandlers {
		registerableHandler(server, appCtx, bundle)
	}
}