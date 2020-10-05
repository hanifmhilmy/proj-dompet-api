package services

import (
	"database/sql"
	"log"

	"github.com/hanifmhilmy/proj-dompet-api/app/domain/model"
	"github.com/hanifmhilmy/proj-dompet-api/app/domain/repository"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type (
	UserServiceInterface interface {
		SaveCreatedUser(tx *sqlx.Tx, data model.SignUpDetails) error
	}

	userService struct {
		repo repository.UserRepositoryInterface
	}
)

func (u *userService) SaveCreatedUser(tx *sqlx.Tx, data model.SignUpDetails) error {
	userID, err := u.repo.FindAccount(data.Username, data.Password)
	if err != nil && errors.Cause(err) != sql.ErrNoRows {
		log.Println("[UserService] query failed: ", err)
		return err
	}
	if userID > 0 {
		err = errors.New("[UserService] this user is already exist")
		return err
	}

	// Save account
	lastID, err := u.repo.SaveAccount(tx, data.Username, data.Password)
	if err != nil {
		log.Println("[UserService] query save account failed: ", err)
		return err
	}

	// Save account details
	err = u.repo.SaveDetail(tx, lastID, data.Name, data.Email)
	if err != nil {
		log.Println("[UserService] query save detail failed: ", err)
		return err
	}

	return nil
}
