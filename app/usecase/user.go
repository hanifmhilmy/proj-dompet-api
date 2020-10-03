package usecase

import (
	"context"

	"github.com/hanifmhilmy/proj-dompet-api/app/domain/repository"
	"github.com/hanifmhilmy/proj-dompet-api/app/domain/services"
)

type (
	UserUsecaseInterface interface {
		Authorization(ctx context.Context) (token string, err error)
	}

	userUsecase struct {
		r repository.UserRepositoryInterface
		s services.UserServiceInterface
	}
)

func NewUserUsecase(repo repository.UserRepositoryInterface, service services.UserServiceInterface) UserUsecaseInterface {
	return &UserUsecase{
		r: repo,
		s: service,
	}
}

func (u UserUsecase) Authorization(ctx context.Context) (token string, err error) {
	return
}
