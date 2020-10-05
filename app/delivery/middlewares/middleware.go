package middlewares

import (
	"log"
	"net/http"
	"strings"

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

// Authorization middleware to check verify the cookie and pass the value to context
func Authorization(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: check the cookie value and then pass the value to context
		bearToken := r.Header.Get("Authorization")
		strArr := strings.Split(bearToken, " ")

		authVal := ""
		if len(strArr) == 2 {
			authVal = strArr[1]
		}

		ctx := r.Context()
		ctx = helpers.SetTokenContext(ctx, authVal)

		r = r.WithContext(ctx)
		h(w, r)
	}
}
