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

type category struct {
	di registry.DIContainer
}

func RegisterCategory(r *httprouter.Router, ctn registry.DIContainer) {
	c := category{
		di: ctn,
	}
	r.HandlerFunc("GET", "/categories", middlewares.Apply(c.GetCategoryList, middlewares.PanicRecoveryMiddleware, middlewares.SetHeaderOptions))
}

func (c category) GetCategoryList(w http.ResponseWriter, r *http.Request) {
	// Set the timeout
	ctx, cancel := context.WithTimeout(r.Context(), time.Millisecond*500)
	defer cancel()
	u := c.di.Resolve(registry.CategoryUsecase).(usecase.CategoryUsecaseInterface)

	categories, err := u.GetCategoryList(ctx)
	if err != nil {
		helpers.JSONResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	result := map[string][]model.CategoryData{
		"categories": categories,
	}
	helpers.JSONResponse(w, http.StatusOK, result)
}
