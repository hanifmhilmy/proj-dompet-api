package redis

import (
	"github.com/gomodule/redigo/redis"
)

type (
	RedisInterface interface {
	}

	client struct {
		redis.Conn
	}
)

func NewClient() {
}
