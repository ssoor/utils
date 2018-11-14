package redis

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

// Pool is
type Pool interface {
	Get() Conn
	Stats() redis.PoolStats
	ActiveCount() int
	IdleCount() int
	Close() error
}

type pool struct {
	*redis.Pool
}

// Get is
func (c *pool) Get() Conn {
	return &conn{ConnWithTimeout: c.Pool.Get().(redis.ConnWithTimeout)}
}

// NewDefaultPool is
func NewDefaultPool(addr string) Pool {
	rawPool := &redis.Pool{
		MaxIdle:     5,
		MaxActive:   100,
		Wait:        true,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", addr)
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}

	return &pool{Pool: rawPool}
}
