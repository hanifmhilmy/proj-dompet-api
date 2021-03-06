package services

import (
	"github.com/hanifmhilmy/proj-dompet-api/app/domain/model"
	"github.com/hanifmhilmy/proj-dompet-api/app/domain/repository"
	"github.com/hanifmhilmy/proj-dompet-api/pkg/database"
	"github.com/pkg/errors"
)

type (
	UserServiceInterface interface {
		SaveCreatedUser(data model.SignUpDetails) error
	}

	userService struct {
		clientDB database.Client
		repo     repository.UserRepositoryInterface
	}
)

func NewUserService(db database.Client, r repository.UserRepositoryInterface) UserServiceInterface {
	return &userService{
		clientDB: db,
		repo:     r,
	}
}

func (u *userService) SaveCreatedUser(data model.SignUpDetails) error {
	tx, err := u.clientDB.Beginx()
	if err != nil {
		err = errors.Wrap(err, "[UserService] begin failed: ")
		return err
	}
	userID, err := u.repo.FindAccount(data.Username, data.Password)
	if err != nil {
		err = errors.Wrap(err, "[UserService] query failed: ")
		return err
	}
	if userID > 0 {
		err = errors.Wrap(err, "[UserService] this user is already exist")
		return err
	}

	// Save account
	lastID, err := u.repo.SaveAccount(tx, data.Username, data.Password)
	if err != nil {
		err = errors.Wrap(err, "[UserService] query save account failed: ")
		return err
	}

	// Save account details
	err = u.repo.SaveDetail(tx, lastID, data.Name, data.Email)
	if err != nil {
		err = errors.Wrap(err, "[UserService] query save detail failed: ")
		return err
	}

	err = tx.Commit()
	if err != nil {
		err = errors.Wrap(err, "[UserService] commit failed ")
		tx.Rollback()
		return err
	}

	return nil
}
