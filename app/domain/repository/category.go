package repository

import (
	"github.com/hanifmhilmy/proj-dompet-api/app/domain/model"
	"github.com/hanifmhilmy/proj-dompet-api/pkg/database"
	"github.com/hanifmhilmy/proj-dompet-api/pkg/redis"
	"github.com/pkg/errors"
)

type (
	categoryRepository struct {
		db    database.Client
		redis *redis.Redigo
	}
)

// NewCategoryRepo initialize dependency category repository
func NewCategoryRepo(c Client) CategoryRepositoryInterface {
	return &categoryRepository{
		db:    c.DB,
		redis: c.Redis,
	}
}

func (cr categoryRepository) GetCategoryList(tx database.Tx, parentID int64) (data []model.CategoryData, err error) {
	q := tx.Rebind(model.QueryCategoryList)
	rows, err := tx.Queryx(q, parentID)
	if err != nil {
		err = errors.Wrap(err, "[CategoryRepository] Fail to get the list")
		return nil, err
	}

	for rows.Next() {
		cat := model.CategoryData{}
		if errScan := rows.Scan(&cat); errScan != nil {
			err = errors.Wrap(err, errScan.Error())
			return nil, err
		}

		data = append(data, cat)
	}
	return
}
