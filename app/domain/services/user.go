package services

import (
	"github.com/hanifmhilmy/proj-dompet-api/app/domain/repository"
)

type (
	UserServiceInterface interface {
	}

	userService struct {
		repo repository.UserRepositoryInterface
	}
)
