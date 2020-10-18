package repository

import (
	"log"
	"time"

	"github.com/hanifmhilmy/proj-dompet-api/app/domain/model"
	"github.com/hanifmhilmy/proj-dompet-api/pkg/database"
	"github.com/hanifmhilmy/proj-dompet-api/pkg/helpers"
	"github.com/pkg/errors"
)

type (
	balanceRepository struct {
		c Client
	}
)

// NewBalanceRepository is to initialize the balance repository
func NewBalanceRepository(client Client) BalanceRepositoryInterface {
	return &balanceRepository{
		c: client,
	}
}

func (b balanceRepository) Save(tx database.Tx, data model.Balance) error {
	timestamp := time.Now().Format(time.RFC3339Nano)
	q := b.c.DB.Rebind(model.QuerySaveBalance)
	_, err := tx.Exec(q, data.UID, data.Values, data.Name, data.Color, model.BalanceStatusActive, timestamp, data.UID, timestamp, data.UID)
	if err != nil {
		err = errors.Wrap(err, "[Balance Repository] error create ")
		return err
	}

	return nil
}

func (b balanceRepository) Get(tx database.Tx, data model.Balance) error {
	return nil
}

func (b balanceRepository) Update(tx database.Tx, data model.Balance, q helpers.String) error {
	query := ""
	if value, ok := helpers.ToString(q); ok {
		query = value
	}
	log.Println(query)
	return nil
}
