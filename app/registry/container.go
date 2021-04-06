package registry

import (
	"log"
	"net/http"

	"github.com/hanifmhilmy/proj-dompet-api/app/domain/repository"
	"github.com/hanifmhilmy/proj-dompet-api/app/domain/services"
	"github.com/hanifmhilmy/proj-dompet-api/app/interface/persistence/memory"
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
	PostgreMainDB = "postgres-db"
	RedigoClient  = "redigo-client"

	// usecases
	UserUsecase     = "user-usecase"
	BalanceUsecase  = "balance-usecase"
	CategoryUsecase = "category-usecase"

	// repository
	UserRepository        = "user-repository"
	BalanceRepository     = "balance-repository"
	BalanceHistRepository = "balance-hist-repository"
	CategoryRepository    = "category-repository"
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

	// init dependency
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
			Name: UserRepository,
			Build: func(ctn di.Container) (interface{}, error) {
				dbClient := ctn.Get(PostgreMainDB).(database.Client)
				redisClient := ctn.Get(RedigoClient).(*redis.Redigo)
				repo := memory.NewUserRepo(memory.Client{
					DB:    dbClient,
					Redis: redisClient,
				})
				return repo, nil
			},
		},
		{
			Name: BalanceRepository,
			Build: func(ctn di.Container) (interface{}, error) {
				dbClient := ctn.Get(PostgreMainDB).(database.Client)
				redisClient := ctn.Get(RedigoClient).(*redis.Redigo)
				repo := memory.NewBalanceRepository(memory.Client{
					DB:    dbClient,
					Redis: redisClient,
				})
				return repo, nil
			},
		},
		{
			Name: BalanceHistRepository,
			Build: func(ctn di.Container) (interface{}, error) {
				dbClient := ctn.Get(PostgreMainDB).(database.Client)
				redisClient := ctn.Get(RedigoClient).(*redis.Redigo)
				repo := memory.NewBalanceHistRepo(memory.Client{
					DB:    dbClient,
					Redis: redisClient,
				})
				return repo, nil
			},
		},
		{
			Name: CategoryRepository,
			Build: func(ctn di.Container) (interface{}, error) {
				dbClient := ctn.Get(PostgreMainDB).(database.Client)
				redisClient := ctn.Get(RedigoClient).(*redis.Redigo)
				repo := memory.NewCategoryRepo(memory.Client{
					DB:    dbClient,
					Redis: redisClient,
				})
				return repo, nil
			},
		},
		{
			Name: UserUsecase,
			Build: func(ctn di.Container) (interface{}, error) {
				repo, _ := ctn.Get(UserRepository).(repository.UserRepositoryInterface)
				dbClient := ctn.Get(PostgreMainDB).(database.Client)
				service := services.NewUserService(dbClient, repo)

				return usecase.NewUserUsecase(repo, service, auth), nil
			},
		},
		{
			Name: BalanceUsecase,
			Build: func(ctn di.Container) (interface{}, error) {
				dbClient := ctn.Get(PostgreMainDB).(database.Client)
				repo, _ := ctn.Get(BalanceRepository).(repository.BalanceRepositoryInterface)
				repoHist, _ := ctn.Get(BalanceHistRepository).(repository.BalanceHistRepositoryInterface)

				service := services.NewBalanceService(dbClient, repo, repoHist)
				return usecase.NewBalanceUsecase(repo, repoHist, service), nil
			},
		},
		{
			Name: CategoryUsecase,
			Build: func(ctn di.Container) (interface{}, error) {
				dbClient := ctn.Get(PostgreMainDB).(database.Client)
				repo := ctn.Get(CategoryRepository).(repository.CategoryRepositoryInterface)
				service := services.NewCategoryService(dbClient, repo)

				return usecase.NewUsecaseCategory(repo, service), nil
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
