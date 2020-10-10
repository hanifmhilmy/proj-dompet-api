package usecase

import (
	"context"
	"crypto/sha1"
	"encoding/base64"
	"time"

	"github.com/hanifmhilmy/proj-dompet-api/app/domain/model"
	"github.com/hanifmhilmy/proj-dompet-api/app/domain/repository"
	"github.com/hanifmhilmy/proj-dompet-api/app/domain/services"
	"github.com/hanifmhilmy/proj-dompet-api/pkg/auth"
	"github.com/hanifmhilmy/proj-dompet-api/pkg/helpers"
	"github.com/pkg/errors"
)

type (
	UserUsecaseInterface interface {
		Authorization(ctx context.Context, uname, password string) (token map[string]string, err error)
		Register(ctx context.Context, details model.SignUpDetails) error
		VerifyRefreshToken(ctx context.Context) (token map[string]string, err error)
		Logout(ctx context.Context) error
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

func (u userUsecase) Authorization(ctx context.Context, uname, password string) (token map[string]string, err error) {
	// Get the correct hashed uname / password, if it is incorrect return error
	// convert the uname and password to bytes
	hasher := sha1.New()
	hasher.Write([]byte(password))
	hashedPass := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	// check the hashed to match with the data in the database
	userID, err := u.repo.FindAccount(uname, hashedPass)
	if err != nil || userID == 0 {
		err = errors.Wrap(err, "[UserUsecase Authorization] User not found")
		return nil, err
	}

	generatedToken, err := u.auth.CreateToken(userID)
	if err != nil {
		err = errors.Wrap(err, "[UserUsecase Authorization] Token Fail")
		return nil, err
	}

	expire := time.Unix(generatedToken.GetTokenRefreshExpire(), 0)
	access := time.Unix(generatedToken.GetTokenExpire(), 0)
	now := time.Now()
	err = u.repo.SetAccessDetails(model.AccessDetails{
		AccUUID:     generatedToken.GetUUIDAccess(),
		RefreshUUID: generatedToken.GetUUIDRefresh(),
		UserID:      userID,
	}, int64(access.Sub(now).Seconds()), int64(expire.Sub(now).Seconds()))
	if err != nil {
		err = errors.Wrap(err, "[UserUsecase Authorization] Fail to set the redis")
		return nil, err
	}
	return generatedToken.GetToken(), nil
}

// Register regist user to the apps
func (u userUsecase) Register(ctx context.Context, details model.SignUpDetails) error {
	hasher := sha1.New()
	hasher.Write([]byte(details.Password))
	details.Password = base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	err := u.serv.SaveCreatedUser(details)
	if err != nil {
		err = errors.Wrap(err, "[UserUsecase Register] SaveCreatedUser")
		return err
	}

	return nil
}

// VerifyRefreshToken verify token used to refresh the current active token so it will not expired as long the refresh token still exist
func (u userUsecase) VerifyRefreshToken(ctx context.Context) (token map[string]string, err error) {
	tokenRefresh, ok := helpers.GetTokenContext(ctx)
	if !ok || tokenRefresh == "" {
		err = errors.Wrap(auth.ErrInvalidToken, "[UserUsecase VerifyRefreshToken]")
		return nil, err
	}

	userID, ok := helpers.GetUserIDContext(ctx)
	if !ok || userID == model.UserNotFound {
		err = errors.Wrap(auth.ErrUserInvalid, "[UserUsecase VerifyRefreshToken] User not found")
		return nil, err
	}

	userUUID, ok := helpers.GetUserUUIDContext(ctx)
	if !ok || userUUID == "" {
		err = errors.Wrap(auth.ErrUserInvalid, "[UserUsecase VerifyRefreshToken] User not found")
		return nil, err
	}

	if err := u.repo.RemoveTokenCache(userUUID); err != nil {
		err = errors.Wrap(auth.ErrInvalidToken, "[UserUsecase VerifyRefreshToken] fail to remove existing token cache")
		return nil, err
	}

	generatedToken, err := u.auth.CreateToken(userID)
	if err != nil {
		err = errors.Wrap(err, "[UserUsecase VerifyRefreshToken] Token Fail")
		return nil, err
	}

	expire := time.Unix(generatedToken.GetTokenRefreshExpire(), 0)
	access := time.Unix(generatedToken.GetTokenExpire(), 0)
	now := time.Now()
	err = u.repo.SetAccessDetails(model.AccessDetails{
		AccUUID:     generatedToken.GetUUIDAccess(),
		RefreshUUID: generatedToken.GetUUIDRefresh(),
		UserID:      userID,
	}, int64(access.Sub(now).Seconds()), int64(expire.Sub(now).Seconds()))
	if err != nil {
		err = errors.Wrap(err, "[UserUsecase VerifyRefreshToken] Fail to set the redis")
		return nil, err
	}
	return generatedToken.GetToken(), nil
}

func (u userUsecase) Logout(ctx context.Context) (err error) {
	userUUID, ok := helpers.GetUserUUIDContext(ctx)
	if !ok || userUUID == "" {
		err = errors.Wrap(auth.ErrUserInvalid, "[UserUsecase VerifyRefreshToken] User not found")
		return err
	}

	if err := u.repo.RemoveTokenCache(userUUID); err != nil {
		err = errors.Wrap(auth.ErrUserInvalid, "[UserUsecase VerifyRefreshToken] fail to destroy the key")
		return err
	}

	return nil
}
