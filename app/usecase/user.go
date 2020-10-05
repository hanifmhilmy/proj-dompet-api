package usecase

import (
	"context"
	"crypto/sha1"
	"encoding/base64"
	"log"

	"github.com/hanifmhilmy/proj-dompet-api/app/domain/repository"
	"github.com/hanifmhilmy/proj-dompet-api/app/domain/services"
	"github.com/hanifmhilmy/proj-dompet-api/pkg/auth"
)

type (
	UserUsecaseInterface interface {
		Authorization(ctx context.Context, uname, password string) (token map[string]string, err error)
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
	// TODO: check the authentication first, if it already login skip this and return the token directly

	// Get the correct hashed uname / password, if it is incorrect return error
	// convert the uname and password to bytes
	hasher := sha1.New()
	hasher.Write([]byte(password))
	hashedPass := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	// check the hashed to match with the data in the database
	userID, err := u.repo.FindAccount(uname, hashedPass)
	if err != nil || userID == 0 {
		log.Println("[UserUsecase Authorization] User not found, err -> ", err)
		return nil, err
	}

	generatedToken, err := u.auth.CreateToken(userID)
	if err != nil {
		log.Println("[UserUsecase Authorization] Token fail, err -> ", err)
		return nil, err
	}

	// TODO: save to redis / db
	return generatedToken.GetToken(), nil
}
