package redis

import (
	"github.com/gomodule/redigo/redis"
)

func Scan(src []interface{}, dest ...interface{}) ([]interface{}, error) {
	return redis.Scan(src, dest...)
}
