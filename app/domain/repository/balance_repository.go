package repository

import (
	"github.com/hanifmhilmy/proj-dompet-api/app/domain/model"
	"github.com/hanifmhilmy/proj-dompet-api/pkg/database"
	"github.com/hanifmhilmy/proj-dompet-api/pkg/helpers"
)

type (
	// BalanceRepositoryInterface interface wrapper for struct balance repo
	BalanceRepositoryInterface interface {
		Create(tx database.Tx, data model.Balance) error
		Get(tx database.Tx, data model.Balance) (balances []model.BalanceData, err error)
		Update(tx database.Tx, data model.Balance, q helpers.String, args ...interface{}) error
	}
)