package rest

import (
	"github.com/hanifmhilmy/proj-dompet-api/app/interface/rest/handlers"
	"github.com/hanifmhilmy/proj-dompet-api/app/registry"
	"github.com/julienschmidt/httprouter"
)

func Apply(ctn registry.DIContainer) *httprouter.Router {
	router := httprouter.New()
	registerHTTPRoute(router, ctn)

	return router
}

func registerHTTPRoute(r *httprouter.Router, ctn registry.DIContainer) {
	handlers.RegisterHealth(r, ctn)

	handlers.RegisterBalance(r, ctn)
	handlers.RegisterCategory(r, ctn)
	handlers.RegisterUser(r, ctn)
}
