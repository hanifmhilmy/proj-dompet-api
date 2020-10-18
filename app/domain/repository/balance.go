package repository

import (
	"log"

	"github.com/hanifmhilmy/proj-dompet-api/app/domain/model"
	"github.com/hanifmhilmy/proj-dompet-api/pkg/helpers"
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

func (b balanceRepository) Save(data model.Balance) error {
	return nil
}

func (b balanceRepository) Get(data model.Balance) error {
	return nil
}

func (b balanceRepository) Update(data model.Balance, q helpers.String) error {
	query := ""
	if value, ok := helpers.ToString(q); ok {
		query = value
	}
	log.Println(query)
	return nil
}
