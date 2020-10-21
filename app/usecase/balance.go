package usecase

import (
	"context"

	"github.com/hanifmhilmy/proj-dompet-api/app/domain/model"
	"github.com/hanifmhilmy/proj-dompet-api/app/domain/repository"
	"github.com/hanifmhilmy/proj-dompet-api/app/domain/services"
	"github.com/pkg/errors"
)

type (
	BalanceUsecaseInterface interface {
		CreateAccountBalance(ctx context.Context, data model.Balance) error
	}
	balanceUsecase struct {
		repoBalance     repository.BalanceRepositoryInterface
		repoBalanceHist repository.BalanceHistRepositoryInterface
		servBalance     services.BalanceServiceInterface
	}
)

// NewBalanceUsecase initiate the balance usecase
func NewBalanceUsecase(
	rBalance repository.BalanceRepositoryInterface,
	rBalanceHist repository.BalanceHistRepositoryInterface,
	serv services.BalanceServiceInterface,
) BalanceUsecaseInterface {
	return &balanceUsecase{
		repoBalance:     rBalance,
		repoBalanceHist: rBalanceHist,
		servBalance:     serv,
	}
}

func (ub balanceUsecase) CreateAccountBalance(ctx context.Context, data model.Balance) error {
	err := ub.servBalance.CreateAccountBalance(data)
	if err != nil {
		err = errors.Wrap(err, "[Usecase Balance] create")
		return err
	}
	return nil
}
