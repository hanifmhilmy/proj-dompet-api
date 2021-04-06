package memory

import (
	"time"

	"github.com/hanifmhilmy/proj-dompet-api/app/domain/model"
	"github.com/hanifmhilmy/proj-dompet-api/app/domain/repository"
	"github.com/hanifmhilmy/proj-dompet-api/pkg/database"
	"github.com/hanifmhilmy/proj-dompet-api/pkg/helpers"
	"github.com/pkg/errors"
)

type (
	balanceRepository struct {
		c Client
	}
)

const (
	queryCreateBalance = `INSERT INTO balance(user_id, last_values, name, color, currency_id, status, create_time, create_by, update_time, update_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	queryGetBalance    = `SELECT name, last_values, color, status, currency_id FROM balance where status = 1 and user_id = $1`
	queryUpdateBalance = `UPDATE balance SET name = $1, color = $2, last_values = $3, status = $4, currency_id = $5, update_time = $6, update_by = $7 WHERE balance_id = $8`
)

// NewBalanceRepository is to initialize the balance repository
func NewBalanceRepository(client Client) repository.BalanceRepositoryInterface {
	return &balanceRepository{
		c: client,
	}
}

func (b balanceRepository) Create(tx database.Tx, data model.Balance) error {
	timestamp := time.Now().Format(time.RFC3339Nano)
	q := b.c.DB.Rebind(queryCreateBalance)
	_, err := tx.Exec(q, data.UID, data.Values, data.Name, data.Color, data.CurrencyID, model.BalanceStatusActive, timestamp, data.UID, timestamp, data.UID)
	if err != nil {
		err = errors.Wrap(err, "[Balance Repository] error create ")
		return err
	}

	return nil
}

func (b balanceRepository) Get(tx database.Tx, data model.Balance) (balances []model.BalanceData, err error) {
	q := tx.Rebind(queryGetBalance)
	rows, err := b.c.DB.Queryx(q, data.UID)
	if err != nil {
		err = errors.Wrap(err, "[Balance Repository] error get ")
		return nil, err
	}
	for rows.Next() {
		data := model.BalanceData{}
		err = rows.Scan(&data)
		if err != nil {
			err = errors.Wrap(err, "[Balance Repository] error scan")
			return nil, err
		}
		balances = append(balances, data)
	}

	return balances, nil
}

func (b balanceRepository) Update(tx database.Tx, data model.Balance, q helpers.String, args ...interface{}) error {
	timestamp := time.Now().Format(time.RFC3339Nano)
	query := tx.Rebind(queryUpdateBalance)
	if value, ok := helpers.ToString(q); ok {
		query = tx.Rebind(value)
	}
	_, err := tx.Exec(query, data.Name, data.Color, data.Values, data.Status, data.CurrencyID, timestamp, data.UID, data.ID)
	if err != nil {
		err = errors.Wrap(err, "[Balance Repository] error exec ")
		return err
	}
	return nil
}
