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
	"github.com/pkg/errors"
)

type (
	UserUsecaseInterface interface {
		Authorization(ctx context.Context, uname, password string) (token map[string]string, err error)
		IsAuthorized(ctx context.Context, tokenString string) (userID int64, err error)
		Register(ctx context.Context, details model.SignUpDetails) error
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

func (u userUsecase) IsAuthorized(ctx context.Context, tokenString string) (userID int64, err error) {
	jwtToken, err := u.auth.VerifyToken(tokenString)
	if err != nil {
		err = errors.Wrap(err, "[UserUsecase IsAuthorized] VerifyToken")
		return
	}

	if jwtToken == nil {
		err = errors.Wrap(errors.New("Unauthorized"), "[UserUsecase IsAuthorized] Token nil")
		return
	}

	if !jwtToken.Valid {
		err = errors.Wrap(errors.New("Unauthorized"), "[UserUsecase IsAuthorized] Token Invalid")
		return
	}

	detail, err := u.auth.ExtractTokenMetadata(jwtToken)
	if err != nil {
		err = errors.Wrap(err, "[UserUsecase IsAuthorized] Extract Token")
		return
	}

	err = u.repo.GetAccessDetails(detail.UUID)
	if err != nil {
		err = errors.Wrap(errors.New("Unauthorized"), "[UserUsecase IsAuthorized] GetAccessDetails")
		return
	}

	return detail.UserID, nil
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
