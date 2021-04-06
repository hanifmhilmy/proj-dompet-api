package repository

import (
	"github.com/hanifmhilmy/proj-dompet-api/app/domain/model"
	"github.com/hanifmhilmy/proj-dompet-api/pkg/database"
)

type (
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