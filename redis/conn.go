package redis

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

// Conn represents a connection to a Redis server.
type Conn interface {
	// Close closes the connection.
	Close() error

	// Err returns a non-nil value when the connection is not usable.
	Err() error

	// Send writes the command to the client's output buffer.
	Send(commandName string, args ...interface{}) error

	// Flush flushes the output buffer to the Redis server.
	Flush() error

	// Do sends a command to the server and returns the received reply.
	Do(commandName string, args ...interface{}) (reply interface{}, err error)

	// Do sends a command to the server and returns the received reply.
	// The timeout overrides the read timeout set when dialing the
	// connection.
	DoWithTimeout(timeout time.Duration, commandName string, args ...interface{}) (reply interface{}, err error)

	// GetPubSubConn is
	GetPubSubConn() PubSubConn

	// Subscribe subscribes the connection to the specified channels.
	Subscribe(channel ...interface{}) error
	// PSubscribe subscribes the connection to the given patterns.
	PSubscribe(channel ...interface{}) error
	// Unsubscribe unsubscribes the connection from the given channels, or from all
	// of them if none is given.
	Unsubscribe(channel ...interface{}) error
	// PUnsubscribe unsubscribes the connection from the given patterns, or from all
	// of them if none is given.
	PUnsubscribe(channel ...interface{}) error

	// Receive receives a single reply from the Redis server
	Receive() (reply interface{}, err error)

	// Receive receives a single reply from the Redis server. The timeout
	// overrides the read timeout set when dialing the connection.
	ReceiveWithTimeout(timeout time.Duration) (reply interface{}, err error)
}

type conn struct {
	redis.ConnWithTimeout
}

func (c *conn) GetPubSubConn() PubSubConn {
	return &pubSubConn{PubSubConn: &redis.PubSubConn{Conn: c.ConnWithTimeout}}
}

func (c *conn) Receive() (reply interface{}, err error) {
	pubsubConn := &redis.PubSubConn{Conn: c.ConnWithTimeout}

	switch v := pubsubConn.Receive().(type) {
	case error:
		return nil, v
	case redis.Message:
		return Message{Channel: v.Channel, Pattern: v.Pattern, Data: v.Data}, nil
	default:
		return v, nil
	}
}
func (c *conn) ReceiveWithTimeout(timeout time.Duration) (reply interface{}, err error) {
	pubsubConn := &redis.PubSubConn{Conn: c.ConnWithTimeout}

	switch v := pubsubConn.ReceiveWithTimeout(timeout).(type) {
	case error:
		return nil, v
	case redis.Message:
		return Message{Channel: v.Channel, Pattern: v.Pattern, Data: v.Data}, nil
	default:
		return v, nil
	}
}

func (c *conn) Subscribe(channel ...interface{}) error {
	pubConn := redis.PubSubConn{Conn: c.ConnWithTimeout}

	return pubConn.Subscribe(channel...)
}

func (c *conn) PSubscribe(channel ...interface{}) error {
	pubConn := redis.PubSubConn{Conn: c.ConnWithTimeout}

	return pubConn.PSubscribe(channel...)
}

func (c *conn) Unsubscribe(channel ...interface{}) error {
	pubConn := redis.PubSubConn{Conn: c.ConnWithTimeout}

	return pubConn.Unsubscribe(channel...)
}

func (c *conn) PUnsubscribe(channel ...interface{}) error {
	pubConn := redis.PubSubConn{Conn: c.ConnWithTimeout}

	return pubConn.Unsubscribe(channel...)
}
