package repository

import (
	"database/sql"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/hanifmhilmy/proj-dompet-api/app/domain/model"
	"github.com/hanifmhilmy/proj-dompet-api/pkg/database"
	"github.com/pkg/errors"
)

type (
	UserRepositoryInterface interface {
		FindAccount(uname, password string) (int64, error)
		FindAccountDetail(userID int64) (*model.Account, error)
		SaveAccount(tx database.Tx, user, password string) (int64, error)
		SaveDetail(tx database.Tx, userID int64, name, email string) error
		GetAccessDetails(userUUID string) error
		SetAccessDetails(details model.AccessDetails, expireAccess, expireRefresh int64) error
		RemoveTokenCache(uuid string) error
	}

	userRepository struct {
		db    database.Client
		redis redis.Conn
	}

	Client struct {
		DB    database.Client
		Redis redis.Conn
	}
)

func NewUserRepo(c Client) UserRepositoryInterface {
	return &userRepository{
		db:    c.DB,
		redis: c.Redis,
	}
}

func (r *userRepository) FindAccount(uname, password string) (int64, error) {
	var userID int64
	q := "select user_id from account where status=1 and username=$1 and password=$2"
	q = r.db.Rebind(q)
	err := r.db.QueryRowx(q, uname, password).Scan(&userID)
	if err != nil && errors.Cause(err) != sql.ErrNoRows {
		err = errors.Wrap(err, "[UserRepository DB] Fail when lookup account, err -> ")
		return userID, err
	}

	return userID, nil
}

func (r *userRepository) FindAccountDetail(userID int64) (*model.Account, error) {
	var ac model.AccountData

	err := r.db.Select(&ac, "select user_id, name, email, create_time, create_by, update_time, update_by from account_detail where user_id=?", userID)
	if err != nil && errors.Cause(err) != sql.ErrNoRows {
		err = errors.Wrap(err, "[UserRepository DB] Fail when lookup account detail, err -> ")
		return nil, err
	}

	return model.NewUser(ac), nil
}

func (r *userRepository) SaveAccount(tx database.Tx, user, password string) (uid int64, err error) {
	// TODO: change status active to pending after implement verification
	q := "insert into account (username, password, status, create_time, create_by, update_time, update_by) values ($1, $2, $3, $4, $5, $6, $7) returning user_id"
	q = tx.Rebind(q)
	currentTime := time.Now().Format(time.RFC3339Nano)
	err = tx.QueryRowx(q, user, password, model.UserStatusActive, currentTime, model.UserActionBySystem, currentTime, model.UserActionBySystem).Scan(&uid)
	if err != nil {
		err = errors.Wrap(err, "[UserRepository DB] Fail to save account data")
		return
	}
	return uid, nil
}

func (r *userRepository) SaveDetail(tx database.Tx, userID int64, name, email string) error {
	// TODO: change status active to pending after implement verification
	q := "insert into account_detail (user_id, name, email, create_time, create_by, update_time, update_by) values ($1, $2, $3, $4, $5, $6, $7)"
	q = tx.Rebind(q)
	currentTime := time.Now().Format(time.RFC3339Nano)
	_, err := tx.Exec(q, userID, name, email, currentTime, model.UserActionBySystem, currentTime, model.UserActionBySystem)
	if err != nil {
		err = errors.Wrap(err, "[UserRepository DB] Fail to save account data")
		return err
	}
	return nil
}

// GetAccessDetails Get the redis value of the access token
func (r *userRepository) GetAccessDetails(userUUID string) error {
	_, err := r.redis.Do("GET", userUUID)
	if err != nil {
		err = errors.Wrap(err, "[UserRepository Redis] Empty Redis ")
		return err
	}
	return nil
}

// SetAccessDetails set the redis value for the granted access token
func (r *userRepository) SetAccessDetails(details model.AccessDetails, expireAccess, expireRefresh int64) error {
	_, err := r.redis.Do("SET", details.AccUUID, details.UserID)
	if err != nil {
		err = errors.Wrap(err, "[UserRepository Redis] Fail Set Access Key")
		return err
	}
	_, err = r.redis.Do("EXPIRE", details.AccUUID, expireAccess)
	if err != nil {
		err = errors.Wrap(err, "[UserRepository Redis] Fail Set Expire Access Key")
		return err
	}

	// Set Refresh Token Cache
	_, err = r.redis.Do("SET", details.RefreshUUID, details.UserID)
	if err != nil {
		err = errors.Wrap(err, "[UserRepository Redis] Fail Set Refresh Key")
		return err
	}
	_, err = r.redis.Do("EXPIRE", details.RefreshUUID, expireRefresh)
	if err != nil {
		err = errors.Wrap(err, "[UserRepository Redis] Fail Set Expire Refresh Key")
		return err
	}

	return nil
}

// RemoveRefreshToken remove old refresh access token
func (r *userRepository) RemoveTokenCache(uuid string) error {
	// Delete Refresh Token Cache
	result, err := r.redis.Do("DEL", uuid)
	if err != nil {
		err = errors.Wrap(err, "[UserRepository Redis] Fail Delete Refresh Key")
		return err
	}

	if result.(int64) == 0 {
		return errors.New("No Key Found")
	}

	return nil
}
