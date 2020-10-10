package delivery

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/hanifmhilmy/proj-dompet-api/app/domain/model"
	"github.com/hanifmhilmy/proj-dompet-api/app/registry"
	"github.com/hanifmhilmy/proj-dompet-api/app/usecase"
	"github.com/hanifmhilmy/proj-dompet-api/pkg/helpers"
)

func (h Handler) Authorization(w http.ResponseWriter, r *http.Request) {
	// Set the timeout
	ctx, cancel := context.WithTimeout(r.Context(), time.Millisecond*500)
	defer cancel()

	var login model.LoginDetails
	u := h.di.Resolve(registry.UserUsecase).(usecase.UserUsecaseInterface)
	err := helpers.ReadJSONBody(r, &login)
	if err != nil {
		log.Println(err)
		helpers.JSONResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	isValid, err := helpers.Validate(login)
	if err != nil && !isValid {
		helpers.JSONResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	token, err := u.Authorization(ctx, login.Username, login.Password)
	if err != nil {
		log.Println(err)
		helpers.JSONResponse(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusOK, token)
}

func (h Handler) Register(w http.ResponseWriter, r *http.Request) {
	// Set the timeout
	ctx, cancel := context.WithTimeout(r.Context(), time.Millisecond*500)
	defer cancel()

	var signup model.SignUpDetails
	err := helpers.ReadJSONBody(r, &signup)
	if err != nil {
		log.Println(err)
		helpers.JSONResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	isValid, err := helpers.Validate(signup)
	if err != nil && !isValid {
		helpers.JSONResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	u := h.di.Resolve(registry.UserUsecase).(usecase.UserUsecaseInterface)
	err = u.Register(ctx, signup)
	if err != nil {
		log.Println(err)
		helpers.JSONResponse(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusOK, model.SignUpSuccess)
}

func (h Handler) Verify(w http.ResponseWriter, r *http.Request) {
	// Set the timeout
	ctx, cancel := context.WithTimeout(r.Context(), time.Millisecond*500)
	defer cancel()

	u := h.di.Resolve(registry.UserUsecase).(usecase.UserUsecaseInterface)
	token, err := u.VerifyRefreshToken(ctx)
	if err != nil {
		log.Println(err)
		helpers.JSONResponse(w, http.StatusUnauthorized, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusOK, token)
}

func (h Handler) Logout(w http.ResponseWriter, r *http.Request) {
	// Set the timeout
	ctx, cancel := context.WithTimeout(r.Context(), time.Millisecond*500)
	defer cancel()

	u := h.di.Resolve(registry.UserUsecase).(usecase.UserUsecaseInterface)
	if err := u.Logout(ctx); err != nil {
		log.Println(err)
		helpers.JSONResponse(w, http.StatusUnauthorized, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusOK, model.LoggedOutSuccess)
}
