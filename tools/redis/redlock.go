package redis

import "github.com/gomodule/redigo/redis"

var LockDefaultTTL = 60

type Redlock struct {
	pool *redis.Pool
	ttl  int
	key  string
}

func NewRedlock(client *Client, key string) *Redlock {
	return &Redlock{
		pool: client.Pool,
		ttl:  LockDefaultTTL,
		key:  key,
	}
}

func (rl *Redlock) Setnx() (bool, error) {
	conn := rl.pool.Get()
	defer conn.Close()

	reply, err := conn.Do("SETNX", rl.key, 1)
	if err != nil {
		return false, err
	}
	if reply.(int64) != 1 {
		return false, nil
	}
	_, err = conn.Do("EXPIRE", rl.key, rl.ttl)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (rl *Redlock) Expire() error {
	conn := rl.pool.Get()
	defer conn.Close()

	_, err := conn.Do("DEL", rl.key)
	return err
}

func (rl *Redlock) SetTtl(t int) {
	rl.ttl = t
}
