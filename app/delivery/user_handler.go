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
	err := helpers.ReadJSONBody(r, &login)
	if err != nil {
		log.Println(err)
		helpers.JSONResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	u := h.di.Resolve(registry.UserUsecase).(usecase.UserUsecaseInterface)
	token, err := u.Authorization(ctx, login.Username, login.Password)
	if err != nil {
		helpers.JSONResponse(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusOK, token)
}
