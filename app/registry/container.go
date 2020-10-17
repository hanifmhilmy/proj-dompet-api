package registry

import (
	"log"
	"net/http"

	"github.com/hanifmhilmy/proj-dompet-api/app/domain/repository"
	"github.com/hanifmhilmy/proj-dompet-api/app/domain/services"
	"github.com/hanifmhilmy/proj-dompet-api/app/usecase"
	"github.com/hanifmhilmy/proj-dompet-api/config"
	"github.com/hanifmhilmy/proj-dompet-api/pkg/auth"
	"github.com/hanifmhilmy/proj-dompet-api/pkg/database"
	"github.com/hanifmhilmy/proj-dompet-api/pkg/redis"
	"github.com/sarulabs/di"
)

// DIContainer interface container for sarulabs di
type DIContainer interface {
	HTTPMiddleware(h http.HandlerFunc) http.HandlerFunc
	Resolve(name string) interface{}
	Clean() error
}

// Container is the default struct to store di Container
type Container struct {
	ctn di.Container
}

const (
	// PostgreMainDB container built for db main connection
	PostgreMainDB   = "postgres-db"
	RedigoClient    = "redigo-client"
	UserUsecase     = "user-usecase"
	CategoryUsecase = "category-usecase"
)

// NewContainer is to init new app container
func NewContainer(conf config.Config) (DIContainer, error) {
	builder, err := di.NewBuilder()
	if err != nil {
		return nil, err
	}

	auth := auth.NewAuth(auth.Options{
		AccessExpire:  conf.Token.AccessExpire,
		RefreshExpire: conf.Token.RefreshExpire,
	})

	db := database.NewDB(conf)
	rdg := redis.New(conf)
	db.Connect([]string{database.DBMain})
	if err := builder.Add([]di.Def{
		{
			Name: PostgreMainDB,
			Build: func(ctn di.Container) (interface{}, error) {
				return db.GetDB(database.DBMain)
			},
		},
		{
			Name: RedigoClient,
			Build: func(ctn di.Container) (interface{}, error) {
				if _, err := rdg.Ping(); err != nil {
					return nil, err
				}
				return rdg, nil
			},
		},
		{
			Name: UserUsecase,
			Build: func(ctn di.Container) (interface{}, error) {
				dbClient := ctn.Get(PostgreMainDB).(database.Client)
				redisClient := ctn.Get(RedigoClient).(*redis.Redigo)
				repo := repository.NewUserRepo(repository.Client{
					DB:    dbClient,
					Redis: redisClient,
				})
				return usecase.NewUserUsecase(repo, services.NewUserService(repo, dbClient, redisClient), auth), nil
			},
		},
		{
			Name: CategoryUsecase,
			Build: func(ctn di.Container) (interface{}, error) {
				dbClient := ctn.Get(PostgreMainDB).(database.Client)
				redisClient := ctn.Get(RedigoClient).(*redis.Redigo)

				repoClient := repository.Client{
					DB:    dbClient,
					Redis: redisClient,
				}
				repo := repository.NewCategoryRepo(repoClient)
				return usecase.NewUsecaseCategory(repo, services.NewCategoryService(repoClient, repo)), nil
			},
		},
	}...); err != nil {
		return nil, err
	}

	return &Container{
		ctn: builder.Build(),
	}, nil
}

// HTTPMiddleware register http middleware function
func (c *Container) HTTPMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return di.HTTPMiddleware(h, c.ctn, func(msg string) {
		log.Println("Captured: ", msg)
	})
}

// Resolve for resolving the function which initialized by the New function
func (c *Container) Resolve(name string) interface{} {
	return c.ctn.Get(name)
}

// Clean for cleaning up the DI
func (c *Container) Clean() error {
	return c.ctn.Clean()
}
