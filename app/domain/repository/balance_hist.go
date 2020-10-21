package repository

import (
	"github.com/hanifmhilmy/proj-dompet-api/app/domain/model"
	"github.com/hanifmhilmy/proj-dompet-api/pkg/database"
)

type (
	balanceHist struct {
		c Client
	}
)

const (
	queryHistSave = `INSERT INTO balance(
		user_id, last_values, name, color, status, create_time, create_by, update_time, update_by
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	queryHistGet    = `SELECT name, last_values, color, status FROM balance where status = 1 and user_id = $1`
	queryHistUpdate = `UPDATE balance SET name = $1, color = $2, last_values = $3, status = $4, update_time = $5, update_by = $6 WHERE balance_id = $7`
)

// NewBalanceHistRepo to initialize the balance history repository dependency
func NewBalanceHistRepo(client Client) BalanceHistRepositoryInterface {
	return &balanceHist{
		c: client,
	}
}

func (bh balanceHist) Create(tx database.Tx, data model.BalanceHist) error {
	return nil
}
