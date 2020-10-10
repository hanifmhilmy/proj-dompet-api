package middlewares

import (
	"log"
	"net/http"
	"strings"

	"github.com/hanifmhilmy/proj-dompet-api/pkg/auth"
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

// IsAuthorized middleware to verify the authorization
func IsAuthorized(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		bearToken := r.Header.Get("Authorization")
		strArr := strings.Split(bearToken, " ")

		if len(strArr) != 2 {
			helpers.JSONResponse(w, http.StatusForbidden, auth.ErrMalformedToken.Error())
			return
		}

		jwtToken, err := auth.VerifyToken(strArr[1], auth.AccessToken)
		if err != nil {
			helpers.JSONResponse(w, http.StatusUnauthorized, auth.ErrInvalidToken.Error())
			return
		}

		if jwtToken == nil || (jwtToken != nil && !jwtToken.Valid) {
			helpers.JSONResponse(w, http.StatusForbidden, auth.ErrInvalidToken.Error())
			return
		}

		detail, err := auth.ExtractTokenMetadata(jwtToken, auth.ClaimUUIDAccess)
		if err != nil || detail == nil {
			helpers.JSONResponse(w, http.StatusForbidden, auth.ErrExtractTokenMetadata.Error())
			return
		}

		ctx = helpers.SetTokenContext(ctx, strArr[1])
		ctx = helpers.SetUserIDContext(ctx, detail.UserID)
		ctx = helpers.SetUserUUIDContext(ctx, detail.UUID)

		r = r.WithContext(ctx)
		h(w, r)
	}
}

// RefreshToken middleware to verify the refresh token in the cookie
func RefreshToken(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		cookie, err := r.Cookie("_RID_Pundi_")
		if err == http.ErrNoCookie {
			log.Println(cookie)
			helpers.JSONResponse(w, http.StatusForbidden, auth.ErrMalformedToken.Error())
			return
		}

		jwtToken, err := auth.VerifyToken(cookie.Value, auth.RefreshToken)
		if err != nil {
			helpers.JSONResponse(w, http.StatusUnauthorized, auth.ErrInvalidToken.Error())
			return
		}

		if jwtToken == nil || (jwtToken != nil && !jwtToken.Valid) {
			helpers.JSONResponse(w, http.StatusForbidden, auth.ErrInvalidToken.Error())
			return
		}

		detail, err := auth.ExtractTokenMetadata(jwtToken, auth.ClaimUUIDRefresh)
		if err != nil || detail == nil {
			helpers.JSONResponse(w, http.StatusForbidden, auth.ErrExtractTokenMetadata.Error())
			return
		}

		ctx = helpers.SetTokenContext(ctx, cookie.Value)
		ctx = helpers.SetUserIDContext(ctx, detail.UserID)
		ctx = helpers.SetUserUUIDContext(ctx, detail.UUID)

		r = r.WithContext(ctx)
		h(w, r)
	}
}
