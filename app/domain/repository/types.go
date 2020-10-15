package repository

import (
	"github.com/hanifmhilmy/proj-dompet-api/app/domain/model"
	"github.com/hanifmhilmy/proj-dompet-api/pkg/database"
	"github.com/hanifmhilmy/proj-dompet-api/pkg/redis"
	"github.com/pkg/errors"
)

// interface declaration
type (
	// CategoryRepositoryInterface interface wrapper for struct category repo
	CategoryRepositoryInterface interface {
		GetCategoryList(tx database.Tx, parentID int64) (data []model.CategoryData, err error)
	}

	// UserRepositoryInterface interface wrapper for struct user repo
	UserRepositoryInterface interface {
		FindAccount(uname, password string) (int64, error)
		FindAccountDetail(userID int64) (*model.AccountData, error)
		SaveAccount(tx database.Tx, user, password string) (int64, error)
		SaveDetail(tx database.Tx, userID int64, name, email string) error
		GetAccessDetails(userUUID string) error
		SetAccessDetails(details model.AccessDetails, expireAccess, expireRefresh int64) error
		RemoveTokenCache(uuid string) error
		SetPasswordResetTokenCache(uid int64, value string) error
		GetPasswordResetTokenCache(uid int64) (string, error)
		RemovePasswordResetTokenCache(uid int64) error
	}
)

// Registry client types
type (
	// Client struct to store the dependency used by repo
	Client struct {
		DB    database.Client
		Redis *redis.Redigo
	}
)

var (
	//ErrKeyNotFound missing key redis cache
	ErrKeyNotFound = errors.New("No Key Found")
)
