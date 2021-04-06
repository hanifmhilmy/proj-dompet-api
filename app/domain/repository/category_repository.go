package repository

import (
	"github.com/hanifmhilmy/proj-dompet-api/app/domain/model"
	"github.com/hanifmhilmy/proj-dompet-api/pkg/database"
)

type (
	// CategoryRepositoryInterface interface wrapper for struct category repo
	CategoryRepositoryInterface interface {
		GetCategoryList(tx database.Tx, parentID int64) (data []model.CategoryData, err error)
	}
)