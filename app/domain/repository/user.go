package repository

import (
	"log"
	"time"

	"github.com/hanifmhilmy/proj-dompet-api/app/domain/model"
	"github.com/jmoiron/sqlx"
)

type (
	UserRepositoryInterface interface {
		FindAccount(uname, password string) (int64, error)
		FindAccountDetail(userID int64) (*model.Account, error)
		SaveAccount(tx *sqlx.Tx, user, password string) (int64, error)
		SaveDetail(tx *sqlx.Tx, userID int64, name, email string) error
	}

	userRepository struct {
		db *sqlx.DB
	}

	Client struct {
		DB *sqlx.DB
	}
)

func NewUserRepo(c Client) UserRepositoryInterface {
	return &userRepository{
		db: c.DB,
	}
}

func (r *userRepository) FindAccount(uname, password string) (int64, error) {
	var userID int64
	q := "select user_id from account where status=1 and username=$1 and password=$2"
	q = r.db.Rebind(q)
	err := r.db.QueryRow(q, uname, password).Scan(&userID)
	if err != nil {
		log.Println("[UserRepository DB] Fail when lookup account, err -> ", err)
		return userID, err
	}

	return userID, nil
}

func (r *userRepository) FindAccountDetail(userID int64) (*model.Account, error) {
	var ac model.AccountData

	err := r.db.Select(&ac, "select user_id, name, email, create_time, create_by, update_time, update_by from account_detail where user_id=?", userID)
	if err != nil {
		log.Println("[UserRepository DB] Fail when lookup account detail, err -> ", err)
		return nil, err
	}

	return model.NewUser(ac), nil
}

func (r *userRepository) SaveAccount(tx *sqlx.Tx, user, password string) (uid int64, err error) {
	// TODO: change status active to pending after implement verification
	q := "insert into account (username, password, status, create_time, create_by, update_time, update_by) values ($1, $2, $3, $4, $5, $6, $7)"
	q = tx.Rebind(q)
	currentTime := time.Now().Format(time.RFC3339Nano)
	result, err := tx.Exec(q, user, password, model.UserStatusActive, currentTime, model.UserActionBySystem, currentTime, model.UserActionBySystem)
	if err != nil {
		log.Println("[UserRepository DB] Fail to save account data")
		return
	}
	userID, err := result.LastInsertId()
	if err != nil {
		log.Println("[UserRepository DB] Fail to retrieve last inserted id, err -> ", err)
		return
	}
	return userID, nil
}

func (r *userRepository) SaveDetail(tx *sqlx.Tx, userID int64, name, email string) error {
	// TODO: change status active to pending after implement verification
	q := "insert into account_detail (user_id, name, email, create_time, create_by, update_time, update_by) values ($1, $2, $3, $4, $5, $6, $7)"
	q = tx.Rebind(q)
	currentTime := time.Now().Format(time.RFC3339Nano)
	_, err := tx.Exec(q, userID, name, email, currentTime, model.UserActionBySystem, currentTime, model.UserActionBySystem)
	if err != nil {
		log.Println("[UserRepository DB] Fail to save account data")
		return err
	}
	return nil
}
