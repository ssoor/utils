package redis

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

// PubSubConn is
type PubSubConn interface {
	Close() error
	Subscribe(channel ...interface{}) error
	PSubscribe(channel ...interface{}) error
	Unsubscribe(channel ...interface{}) error
	Receive() interface{}
	ReceiveWithTimeout(timeout time.Duration) interface{}
}

type pubSubConn struct {
	*redis.PubSubConn
}

func (c *pubSubConn) Receive() interface{} {
	switch v := c.PubSubConn.Receive().(type) {
	case redis.Message:
		return Message{Channel: v.Channel, Pattern: v.Pattern, Data: v.Data}
	default:
		return v
	}
}
