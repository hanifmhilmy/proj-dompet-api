package delivery

import (
	"net/http"

	"github.com/hanifmhilmy/proj-dompet-api/app/registry"
)

type Handler struct {
	di registry.DIContainer
}

func NewHandler(ctn registry.DIContainer) *Handler {
	return &Handler{
		di: ctn,
	}
}

func (h Handler) Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}
