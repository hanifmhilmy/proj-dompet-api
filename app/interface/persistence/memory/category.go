package memory

import (
	"github.com/hanifmhilmy/proj-dompet-api/app/domain/model"
	"github.com/hanifmhilmy/proj-dompet-api/app/domain/repository"
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

const (
	queryCategoryList = `SELECT category_id, value, parent_id, create_time, create_by, update_time, update_by from category where status = 1 and parent_id = $1`
)

// NewCategoryRepo initialize dependency category repository
func NewCategoryRepo(c Client) repository.CategoryRepositoryInterface {
	return &categoryRepository{
		db:    c.DB,
		redis: c.Redis,
	}
}

func (cr categoryRepository) GetCategoryList(tx database.Tx, parentID int64) (data []model.CategoryData, err error) {
	q := tx.Rebind(queryCategoryList)
	rows, err := tx.Queryx(q, parentID)
	if err != nil {
		err = errors.Wrap(err, "[CategoryRepository] Fail to get the list")
		return nil, err
	}

	for rows.Next() {
		cat := model.CategoryData{}
		if errScan := rows.StructScan(&cat); errScan != nil {
			err = errors.Wrap(errScan, "[CategoryRepository]")
			return nil, err
		}

		data = append(data, cat)
	}
	return
}
