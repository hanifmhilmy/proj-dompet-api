package handlers

import (
	"net/http"

	"github.com/hanifmhilmy/proj-dompet-api/app/interface/rest/middlewares"
	"github.com/hanifmhilmy/proj-dompet-api/app/registry"
	"github.com/julienschmidt/httprouter"
)

type health struct {
	di registry.DIContainer
}

func RegisterHealth(r *httprouter.Router, ctn registry.DIContainer) {
	h := health{
		di: ctn,
	}

	r.HandlerFunc("GET", "/ping", middlewares.Apply(h.Ping, middlewares.PanicRecoveryMiddleware, middlewares.SetHeaderOptions))
}

func (h health) Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}
