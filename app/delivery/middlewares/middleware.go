package middlewares

import (
	"log"
	"net/http"
)

// PanicRecoveryMiddleware handles the panic in the handlers.
func PanicRecoveryMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				// log the error
				log.Println(rec)

				// write the error response
				w.WriteHeader(500)
				w.Header().Set("Content-Type", "application/json")
				w.Write([]byte("Internal Server Error"))
			}
		}()

		h(w, r)
	}
}
