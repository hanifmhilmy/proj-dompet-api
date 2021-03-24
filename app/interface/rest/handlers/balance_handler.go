package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/hanifmhilmy/proj-dompet-api/app/domain/model"
	"github.com/hanifmhilmy/proj-dompet-api/app/interface/rest/middlewares"
	"github.com/hanifmhilmy/proj-dompet-api/app/registry"
	"github.com/hanifmhilmy/proj-dompet-api/app/usecase"
	"github.com/hanifmhilmy/proj-dompet-api/pkg/helpers"
	"github.com/julienschmidt/httprouter"
)

type balance struct {
	di registry.DIContainer
}

func RegisterBalance(r *httprouter.Router, ctn registry.DIContainer) {
	b := balance{
		di: ctn,
	}

	r.HandlerFunc("POST", "/balance", middlewares.Apply(b.CreateAccountBalance, middlewares.PanicRecoveryMiddleware, middlewares.SetHeaderOptions, middlewares.IsAuthorized))
}

func (b balance) CreateAccountBalance(w http.ResponseWriter, r *http.Request) {
	// Set the timeout
	ctx, cancel := context.WithTimeout(r.Context(), time.Millisecond*500)
	defer cancel()

	userID, ok := helpers.GetUserIDContext(ctx)
	if !ok || userID == model.UserNotFound {
		helpers.CustomJSONResponse(w, http.StatusUnauthorized, "err.context", []string{model.ErrNotLogin.Error()}, nil)
		return
	}

	var data model.Balance
	err := helpers.ReadJSONBody(r, &data)
	if err != nil {
		helpers.CustomJSONResponse(w, http.StatusBadRequest, "err.body", []string{err.Error()}, nil)
		return
	}

	isValid, err := helpers.Validate(data)
	if err != nil && !isValid {
		helpers.JSONResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	data.UID = userID

	u := b.di.Resolve(registry.BalanceUsecase).(usecase.BalanceUsecaseInterface)
	err = u.CreateAccountBalance(ctx, data)
	if err != nil {
		helpers.CustomJSONResponse(w, http.StatusInternalServerError, "err.internal", []string{err.Error()}, nil)
		return
	}

	helpers.CustomJSONResponse(w, http.StatusOK, "", []string{"Success"}, nil)
}
