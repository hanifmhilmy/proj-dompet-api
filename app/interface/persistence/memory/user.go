package memory

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/hanifmhilmy/proj-dompet-api/app/domain/model"
	"github.com/hanifmhilmy/proj-dompet-api/app/domain/repository"
	"github.com/hanifmhilmy/proj-dompet-api/pkg/database"
	"github.com/hanifmhilmy/proj-dompet-api/pkg/redis"
	"github.com/pkg/errors"
)

type (
	userRepository struct {
		db    database.Client
		redis *redis.Redigo
	}
)

// NewUserRepo initialize user repository
func NewUserRepo(c Client) repository.UserRepositoryInterface {
	return &userRepository{
		db:    c.DB,
		redis: c.Redis,
	}
}

// FindAccount to look up a user account in the database
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

// FindAccountDetail to look up a user account details in the database
func (r *userRepository) FindAccountDetail(userID int64) (*model.AccountData, error) {
	var ac model.AccountData

	err := r.db.Select(&ac, "select user_id, name, email, create_time, create_by, update_time, update_by from account_detail where user_id=?", userID)
	if err != nil && errors.Cause(err) != sql.ErrNoRows {
		err = errors.Wrap(err, "[UserRepository DB] Fail when lookup account detail, err -> ")
		return nil, err
	}

	return &ac, nil
}

// SaveAccount to save new user account to the database
func (r *userRepository) SaveAccount(tx database.Tx, user, password string) (uid int64, err error) {
	// TODO: change status active to pending after implement verification
	q := "insert into account (username, password, status, create_time, create_by, update_time, update_by) values ($1, $2, $3, $4, $5, $6, $7) returning user_id"
	q = tx.Rebind(q)
	currentTime := time.Now().Format(time.RFC3339Nano)
	err = tx.QueryRowx(q, user, password, model.UserStatusActive, currentTime, model.ActionBySystem, currentTime, model.ActionBySystem).Scan(&uid)
	if err != nil {
		err = errors.Wrap(err, "[UserRepository DB] Fail to save account data")
		return
	}
	return uid, nil
}

// SaveAccount to save new user account details to the database
func (r *userRepository) SaveDetail(tx database.Tx, userID int64, name, email string) error {
	// TODO: change status active to pending after implement verification
	q := "insert into account_detail (user_id, name, email, create_time, create_by, update_time, update_by) values ($1, $2, $3, $4, $5, $6, $7)"
	q = tx.Rebind(q)
	currentTime := time.Now().Format(time.RFC3339Nano)
	_, err := tx.Exec(q, userID, name, email, currentTime, model.ActionBySystem, currentTime, model.ActionBySystem)
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

// RemoveRefreshToken remove expiry token
func (r *userRepository) RemoveTokenCache(uuid string) error {
	// Delete Refresh Token Cache
	result, err := r.redis.Do("DEL", uuid)
	if err != nil {
		err = errors.Wrap(err, "[UserRepository Redis] Fail to delete token key")
		return err
	}

	if result.(int64) == 0 {
		return ErrKeyNotFound
	}

	return nil
}

// SetPasswordResetTokenCache set the redis password reset token
func (r *userRepository) SetPasswordResetTokenCache(uid int64, value string) error {
	key := fmt.Sprintf(model.RedisResetPassKey, uid)
	_, err := r.redis.Do("SET", key, value)
	if err != nil {
		err = errors.Wrap(err, "[UserRepository Redis] Fail set reset pass token")
		return err
	}

	_, err = r.redis.Do("EXPIRE", key, 300)
	if err != nil {
		err = errors.Wrap(err, "[UserRepository Redis] Fail set expire reset pass token")
		return err
	}

	return nil
}

// GetPasswordResetTokenCache get the password reset token cache
func (r *userRepository) GetPasswordResetTokenCache(uid int64) (string, error) {
	s, err := r.redis.GetString(fmt.Sprintf(model.RedisResetPassKey, uid))
	if err != nil {
		err = errors.Wrap(err, "[UserRepository Redis] Fail retrieve reset token")
		return s, err
	}
	return s, nil
}

// RemovePasswordResetTokenCache get the password reset token cache
func (r *userRepository) RemovePasswordResetTokenCache(uid int64) error {
	rs, err := r.redis.Do(fmt.Sprintf(model.RedisResetPassKey, uid))
	if err != nil {
		err = errors.Wrap(err, "[UserRepository Redis] Fail remove reset token")
		return err
	}

	if rs.(int64) == 0 {
		return ErrKeyNotFound
	}

	return nil
}
