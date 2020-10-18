package model

import (
	"time"

	"github.com/pkg/errors"
)

type (
	// CategoryData model struct
	CategoryData struct {
		ID          int64     `db:"category_id" json:"id"`
		Value       string    `db:"value" json:"value"`
		Status      int       `db:"status" json:"status"`
		ParentID    int64     `db:"parent_id" json:"parent_id"`
		CreatedTime time.Time `db:"create_time" json:"created_time"`
		CreatedBy   int64     `db:"create_by" json:"created_by"`
		UpdateTime  time.Time `db:"update_time" json:"update_time"`
		UpdateBy    int64     `db:"update_by" json:"update_by"`

		Children []CategoryData
	}
)

var (
	ErrMissingCategory = errors.New("Category Not Found")
	ErrInvalidCategory = errors.New("Invalid Category")
)
