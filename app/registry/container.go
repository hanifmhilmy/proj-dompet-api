package registry

import (
	"log"
	"net/http"

	"github.com/hanifmhilmy/proj-dompet-api/config"
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

// NewContainer is to init new app container
func NewContainer(conf config.Config) (DIContainer, error) {
	builder, err := di.NewBuilder()
	if err != nil {
		return nil, err
	}

	if err := builder.Add([]di.Def{}...); err != nil {
		return nil, err
	}

	return &Container{
		ctn: builder.Build(),
	}, nil
}

// HTTPMiddleware register htt pmiddleware function
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
