package handlers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/hanifmhilmy/proj-dompet-api/app/domain/model"
	"github.com/hanifmhilmy/proj-dompet-api/app/interface/rest/middlewares"
	"github.com/hanifmhilmy/proj-dompet-api/app/registry"
	"github.com/hanifmhilmy/proj-dompet-api/app/usecase"
	"github.com/hanifmhilmy/proj-dompet-api/pkg/helpers"
	"github.com/julienschmidt/httprouter"
)

type user struct {
	di registry.DIContainer
}

func RegisterUser(r *httprouter.Router, ctn registry.DIContainer) {
	u := user{
		di: ctn,
	}

	r.HandlerFunc("POST", "/auth", middlewares.Apply(u.Authorization, middlewares.PanicRecoveryMiddleware, middlewares.SetHeaderOptions))
	r.HandlerFunc("PATCH", "/auth", middlewares.Apply(u.Logout, middlewares.PanicRecoveryMiddleware, middlewares.SetHeaderOptions, middlewares.IsAuthorized))
	r.HandlerFunc("GET", "/auth", middlewares.Apply(u.Verify, middlewares.PanicRecoveryMiddleware, middlewares.SetHeaderOptions, middlewares.RefreshToken))
	r.HandlerFunc("POST", "/register", middlewares.Apply(u.Register, middlewares.PanicRecoveryMiddleware, middlewares.SetHeaderOptions))
}

func (u user) Authorization(w http.ResponseWriter, r *http.Request) {
	// Set the timeout
	ctx, cancel := context.WithTimeout(r.Context(), time.Millisecond*500)
	defer cancel()

	var login model.LoginDetails
	uc := u.di.Resolve(registry.UserUsecase).(usecase.UserUsecaseInterface)
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

	token, err := uc.Authorization(ctx, login.Username, login.Password)
	if err != nil {
		log.Println(err)
		helpers.JSONResponse(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusOK, token)
}

func (u user) Register(w http.ResponseWriter, r *http.Request) {
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

	uc := u.di.Resolve(registry.UserUsecase).(usecase.UserUsecaseInterface)
	err = uc.Register(ctx, signup)
	if err != nil {
		log.Println(err)
		helpers.JSONResponse(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusOK, model.SignUpSuccess)
}

func (u user) Verify(w http.ResponseWriter, r *http.Request) {
	// Set the timeout
	ctx, cancel := context.WithTimeout(r.Context(), time.Millisecond*500)
	defer cancel()

	uc := u.di.Resolve(registry.UserUsecase).(usecase.UserUsecaseInterface)
	token, err := uc.VerifyRefreshToken(ctx)
	if err != nil {
		log.Println(err)
		helpers.JSONResponse(w, http.StatusUnauthorized, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusOK, token)
}

func (u user) Logout(w http.ResponseWriter, r *http.Request) {
	// Set the timeout
	ctx, cancel := context.WithTimeout(r.Context(), time.Millisecond*500)
	defer cancel()

	uc := u.di.Resolve(registry.UserUsecase).(usecase.UserUsecaseInterface)
	if err := uc.Logout(ctx); err != nil {
		log.Println(err)
		helpers.JSONResponse(w, http.StatusUnauthorized, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusOK, model.LoggedOutSuccess)
}
