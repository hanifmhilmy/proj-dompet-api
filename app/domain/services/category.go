package services

import (
	"github.com/hanifmhilmy/proj-dompet-api/app/domain/model"
	"github.com/hanifmhilmy/proj-dompet-api/app/domain/repository"
	"github.com/hanifmhilmy/proj-dompet-api/pkg/database"
	"github.com/pkg/errors"
)

type (
	CategoryServiceInterface interface {
		GetCompleteCategoryList() ([]model.CategoryData, error)
	}

	categoryService struct {
		clientDB database.Client
		repo     repository.CategoryRepositoryInterface
	}
)

func NewCategoryService(c repository.Client, r repository.CategoryRepositoryInterface) CategoryServiceInterface {
	return &categoryService{
		clientDB: c.DB,
		repo:     r,
	}
}

func (c categoryService) GetCompleteCategoryList() ([]model.CategoryData, error) {
	tx, err := c.clientDB.Beginx()
	if err != nil {
		err = errors.Wrap(err, "[CategoryService] begin failed: ")
		return nil, err
	}
	categoryList, err := c.repo.GetCategoryList(tx, model.ParentCategory)
	if err != nil {
		err = errors.Wrap(err, "[CategoryService] failed to get parent category list")
		return categoryList, err
	}

	for key := range categoryList {
		parent := categoryList[key]
		children, err := c.repo.GetCategoryList(tx, parent.ParentID)
		if err != nil {
			err = errors.Wrap(err, "[CategoryService] failed to get child category list")
			return categoryList, err
		}

		parent.Children = append(parent.Children, children...)
	}

	return categoryList, nil
}
