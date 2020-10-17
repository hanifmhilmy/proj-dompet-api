package usecase

import (
	"context"

	"github.com/hanifmhilmy/proj-dompet-api/app/domain/model"
	"github.com/hanifmhilmy/proj-dompet-api/app/domain/repository"
	"github.com/hanifmhilmy/proj-dompet-api/app/domain/services"
	"github.com/pkg/errors"
)

type (
	CategoryUsecaseInterface interface {
		GetCategoryList(ctx context.Context) (categories []model.CategoryData, err error)
	}

	categoryUsecase struct {
		repo repository.CategoryRepositoryInterface
		serv services.CategoryServiceInterface
	}
)

func NewUsecaseCategory(repo repository.CategoryRepositoryInterface, service services.CategoryServiceInterface) CategoryUsecaseInterface {
	return &categoryUsecase{
		repo: repo,
		serv: service,
	}
}

func (c categoryUsecase) GetCategoryList(ctx context.Context) (categories []model.CategoryData, err error) {
	categories, err = c.serv.GetCompleteCategoryList()
	if err != nil {
		err = errors.Wrap(err, "[Category Usecase] error on get the list")
	}

	return categories, nil
}
