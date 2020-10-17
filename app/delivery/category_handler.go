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

func (h Handler) GetCategoryList(w http.ResponseWriter, r *http.Request) {
	// Set the timeout
	ctx, cancel := context.WithTimeout(r.Context(), time.Millisecond*500)
	defer cancel()
	u := h.di.Resolve(registry.CategoryUsecase).(usecase.CategoryUsecaseInterface)

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
