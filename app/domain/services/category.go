package services

import (
	"github.com/hanifmhilmy/proj-dompet-api/app/domain/model"
	"github.com/hanifmhilmy/proj-dompet-api/app/domain/repository"
	"github.com/hanifmhilmy/proj-dompet-api/pkg/database"
	"github.com/hanifmhilmy/proj-dompet-api/pkg/redis"
	"github.com/pkg/errors"
)

type (
	CategoryServiceInterface interface {
	}

	categoryService struct {
		clientDB    database.Client
		clientRedis *redis.Redigo
		repo        repository.CategoryRepositoryInterface
	}
)

func NewCategoryService(c repository.Client, r repository.CategoryRepositoryInterface) CategoryServiceInterface {
	return &categoryService{
		clientDB:    c.DB,
		clientRedis: c.Redis,
		repo:        r,
	}
}

func (c categoryService) GetCompleteCategoryList() ([]model.CategoryData, error) {
	categoryList := []model.CategoryData{}
	tx, err := c.clientDB.Beginx()
	if err != nil {
		err = errors.Wrap(err, "[CategoryService] begin failed: ")
		return categoryList, err
	}
	parentList, err := c.repo.GetCategoryList(tx, model.ParentCategory)
	if err != nil {
		err = errors.Wrap(err, "[CategoryService] failed to get parent category list")
		return categoryList, err
	}

	for key := range parentList {
		parent := parentList[key]
		children, err := c.repo.GetCategoryList(tx, parent.ParentID)
		if err != nil {
			err = errors.Wrap(err, "[CategoryService] failed to get child category list")
			return categoryList, err
		}

		parent.Children = append(parent.Children, children...)
	}

	return categoryList, nil
}
