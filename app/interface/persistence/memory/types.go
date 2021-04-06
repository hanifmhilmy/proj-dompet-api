package memory

import (
	"github.com/hanifmhilmy/proj-dompet-api/pkg/database"
	"github.com/hanifmhilmy/proj-dompet-api/pkg/redis"
	"github.com/pkg/errors"
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
	// ErrKeyNotFound missing key redis cache
	ErrKeyNotFound = errors.New("No Key Found")
)
