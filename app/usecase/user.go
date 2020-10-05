package usecase

import (
	"context"

	"github.com/hanifmhilmy/proj-dompet-api/app/domain/repository"
	"github.com/hanifmhilmy/proj-dompet-api/app/domain/services"
	"github.com/hanifmhilmy/proj-dompet-api/pkg/auth"
)

type (
	UserUsecaseInterface interface {
		Authorization(ctx context.Context, uname, password string) (token string, err error)
	}

	userUsecase struct {
		repo repository.UserRepositoryInterface
		serv services.UserServiceInterface
		auth auth.AuthInterface
	}
)

func NewUserUsecase(repo repository.UserRepositoryInterface, service services.UserServiceInterface, auth auth.AuthInterface) UserUsecaseInterface {
	return &userUsecase{
		repo: repo,
		serv: service,
		auth: auth,
	}
}

func (u userUsecase) Authorization(ctx context.Context, uname, password string) (token string, err error) {
	return
}
