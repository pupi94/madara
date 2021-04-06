package redis

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

// cache 默认有效时间
var DefaultTTL = 3 * 24 * time.Hour

type Config struct {
	Namespace string
}

type Client struct {
	Pool   *redis.Pool
	config *Config
}

func NewClient(host string, port int, cfg *Config) *Client {
	pool := &redis.Pool{
		MaxIdle:     15,
		MaxActive:   100,
		IdleTimeout: time.Second * 30,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}

	return &Client{Pool: pool, config: cfg}
}

func (c *Client) Set(key string, value interface{}, expires ...time.Duration) (interface{}, error) {
	conn := c.Pool.Get()
	defer conn.Close()
	return conn.Do("set", c.namespaceKey(key), value, "ex", c.expireTime(expires))
}

func (c *Client) SAdd(key string, fields []interface{}) *Response {
	conn := c.Pool.Get()
	defer conn.Close()
	return NewResponse(conn.Do("sadd", redis.Args{}.Add(c.namespaceKey(key)).AddFlat(fields)...))
}

func (c *Client) LAdd(key string, fields []interface{}) *Response {
	conn := c.Pool.Get()
	defer conn.Close()
	return NewResponse(conn.Do("lpush", redis.Args{}.Add(c.namespaceKey(key)).AddFlat(fields)...))
}

func (c *Client) LRem(key string, count, value interface{}) *Response {
	conn := c.Pool.Get()
	defer conn.Close()

	return NewResponse(conn.Do("lrem", c.namespaceKey(key), count, value))
}

func (c *Client) Exec(fn func(conn redis.Conn) *Response) *Response {
	conn := c.Pool.Get()
	defer conn.Close()
	return fn(conn)
}

//多键删除
func (c *Client) Delete(keys ...string) *Response {
	conn := c.Pool.Get()
	defer conn.Close()

	var fixedKeys []string
	for _, k := range keys {
		fixedKeys = append(fixedKeys, c.namespaceKey(k))
	}
	return NewResponse(conn.Do("DEL", redis.Args{}.AddFlat(keys)...))
}

func (c *Client) expireTime(expires []time.Duration) time.Duration {
	if len(expires) == 0 {
		return DefaultTTL
	}
	return expires[0]
}

func (c *Client) namespaceKey(key string) string {
	if c.config.Namespace == "" {
		return key
	}
	return fmt.Sprintf("%s:%s", c.config.Namespace, key)
}

type Response struct {
	reply interface{}
	err   error
}

func NewResponse(reply interface{}, err error) *Response {
	return &Response{reply: reply, err: err}
}

var LockDefaultTTL = 60

type Redlock struct {
	pool *redis.Pool
	Ttl  int
	Key  string
}

func NewRedlock(client *Client, key string) *Redlock {
	return &Redlock{
		pool: client.Pool,
		Ttl:  LockDefaultTTL,
		Key:  key,
	}
}

func (rl *Redlock) Setnx() (bool, error) {
	conn := rl.pool.Get()
	defer conn.Close()

	reply, err := conn.Do("SETNX", rl.Key, 1)
	if err != nil {
		return false, err
	}
	if reply.(int64) != 1 {
		return false, nil
	}
	_, err = conn.Do("EXPIRE", rl.Key, rl.Ttl)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (rl *Redlock) Expire() error {
	conn := rl.pool.Get()
	defer conn.Close()

	_, err := conn.Do("DEL", rl.Key)
	return err
}
