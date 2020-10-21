package services

import (
	"github.com/hanifmhilmy/proj-dompet-api/app/domain/model"
	"github.com/hanifmhilmy/proj-dompet-api/app/domain/repository"
	"github.com/hanifmhilmy/proj-dompet-api/pkg/database"
	"github.com/pkg/errors"
)

type (
	// BalanceServiceInterface interface service wrapper
	BalanceServiceInterface interface {
		CreateAccountBalance(data model.Balance) error
	}

	balanceService struct {
		clientDB        database.Client
		balanceRepo     repository.BalanceRepositoryInterface
		balanceHistRepo repository.BalanceHistRepositoryInterface
	}
)

// NewBalanceService is to initiate the service dependency related with balance and the balance history
func NewBalanceService(c repository.Client, balance repository.BalanceRepositoryInterface, balanceHist repository.BalanceHistRepositoryInterface) BalanceServiceInterface {
	return &balanceService{
		clientDB:        c.DB,
		balanceRepo:     balance,
		balanceHistRepo: balanceHist,
	}
}

// CreateAccountBalance create the user account balance
func (s balanceService) CreateAccountBalance(data model.Balance) error {
	tx, err := s.clientDB.Beginx()
	if err != nil {
		err = errors.Wrap(err, "[Balance Service] error beginx")
		return err
	}
	err = s.balanceRepo.Create(tx, data)
	if err != nil {
		err = errors.Wrap(err, "[Balance Service]")
		return err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		err = errors.Wrap(err, "[Balance Service]")
		return err
	}

	return nil
}
