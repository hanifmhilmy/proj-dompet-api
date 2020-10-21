package delivery

import (
	"context"
	"net/http"
	"time"

	"github.com/hanifmhilmy/proj-dompet-api/app/domain/model"
	"github.com/hanifmhilmy/proj-dompet-api/app/registry"
	"github.com/hanifmhilmy/proj-dompet-api/app/usecase"
	"github.com/hanifmhilmy/proj-dompet-api/pkg/helpers"
)

func (h Handler) CreateAccountBalance(w http.ResponseWriter, r *http.Request) {
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

	data.UID = userID

	u := h.di.Resolve(registry.BalanceUsecase).(usecase.BalanceUsecaseInterface)
	err = u.CreateAccountBalance(ctx, data)
	if err != nil {
		helpers.CustomJSONResponse(w, http.StatusInternalServerError, "err.internal", []string{err.Error()}, nil)
		return
	}

	helpers.CustomJSONResponse(w, http.StatusOK, "", []string{"Success"}, nil)
}
