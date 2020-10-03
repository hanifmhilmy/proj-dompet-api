package middlewares

import (
	"log"
	"net/http"

	"github.com/hanifmhilmy/proj-dompet-api/pkg/helpers"
)

// SetHeaderOptions set default header options for each request made by the apps
func SetHeaderOptions(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); helpers.CheckWhitelistedOrigin(origin) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		} else {
			w.Header().Set("Access-Control-Allow-Origin", "proj-dompet")
		}
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PATCH,OPTION")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Cookie, Source-Type, Origin, Content-Filename")
		h(w, r)
	}
}

// PanicRecoveryMiddleware handles the panic in the handlers.
func PanicRecoveryMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				// log the error
				log.Println(rec)

				// write the error response
				helpers.JSONResponse(w, http.StatusInternalServerError, rec)
			}
		}()
		h(w, r)
	}
}
