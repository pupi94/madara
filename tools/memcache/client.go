package memcache

import (
	"fmt"
	"strconv"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
)

/*type Client interface {
	Get(key string) (bool, []byte, error)
	GetMulti(keys []string) (map[string][]byte, error)
	Remove(key string) (bool, error)
	Set(key string, value []byte, ttl time.Duration) error
	Increment(key string, delta int64) (uint64, error)
	TryLockOnce(key string) (bool, error)
	Unlock(key string) (bool, error)
}
*/

const (
	// DefaultTimeout is the default socket read/write timeout.
	DefaultTimeout = time.Second

	// DefaultMaxIdleConns is the default maximum number of idle connections
	// kept for any single address.
	DefaultMaxIdleConn = 2
)

type Client struct {
	namespace string
	client    *memcache.Client
}

type Config struct {
	Timeout     time.Duration
	MaxIdleConn int
	Namespace   string
}

func NewClient(servers []string, config *Config) *Client {
	client := memcache.New(servers...)

	if config.MaxIdleConn <= 0 {
		client.MaxIdleConns = DefaultMaxIdleConn
	} else {
		client.MaxIdleConns = config.MaxIdleConn
	}

	if config.Timeout <= 0 {
		client.Timeout = DefaultTimeout
	} else {
		client.Timeout = config.Timeout
	}
	return &Client{client: client, namespace: config.Namespace}
}

func (c *Client) Get(key string) ([]byte, error) {
	item, err := c.client.Get(c.namespaceKey(key))

	if err != nil {
		if err == memcache.ErrCacheMiss {
			return nil, nil
		}
		return nil, err
	}
	return item.Value, nil
}

func (c *Client) GetMulti(keys []string) (map[string][]byte, error) {
	var newKeys = make([]string, len(keys))
	var keyMap = make(map[string]string, len(keys))

	for _, k := range keys {
		newKey := c.namespaceKey(k)
		newKeys = append(newKeys, newKey)
		keyMap[newKey] = k
	}

	result := make(map[string][]byte)
	itemMap, err := c.client.GetMulti(newKeys)
	if err != nil {
		return nil, err
	}
	for k, v := range itemMap {
		result[keyMap[k]] = v.Value
	}
	return result, nil
}

func (c *Client) Remove(key string) (bool, error) {
	err := c.client.Delete(c.namespaceKey(key))

	if err == nil {
		return true, nil
	}
	if err == memcache.ErrCacheMiss {
		return false, nil
	}
	return false, err
}

func (c *Client) Set(key string, value []byte, ttl time.Duration) error {
	if ttl > 30*24*time.Hour {
		ttl = 30 * 24 * time.Hour
	}
	item := &memcache.Item{
		Key:        c.namespaceKey(key),
		Value:      value,
		Expiration: int32(ttl.Seconds()),
	}
	if err := c.client.Set(item); err != nil {
		return err
	}
	return nil
}

func (c *Client) Increment(key string, delta int64) (uint64, error) {
	var f func(string, uint64) (uint64, error)
	if delta > 0 {
		f = c.client.Increment
	} else {
		f = c.client.Decrement
		delta *= -1
	}

	realkey := c.namespaceKey(key)

	newv, err := f(realkey, uint64(delta))
	if err == nil || err != memcache.ErrCacheMiss {
		return newv, err
	}

	//try to add
	err = c.client.Add(&memcache.Item{
		Key:   realkey,
		Value: []byte(strconv.FormatInt(delta, 10)),
	})
	if err == nil {
		return uint64(delta), nil
	} else if err != memcache.ErrNotStored {
		return 0, err
	}

	//if add returns ErrNotStored(key exists), try increment again
	return f(realkey, uint64(delta))
}

func (c *Client) TryLockOnce(key string) (bool, error) {
	err := c.client.Add(&memcache.Item{
		Key: c.namespaceKey(key),
	})

	if err == nil {
		return true, nil
	} else if err == memcache.ErrNotStored {
		return false, nil
	} else {
		return false, err
	}
}

func (c *Client) namespaceKey(key string) string {
	if c.namespace == "" {
		return key
	}
	return fmt.Sprintf("%s:%s", c.namespace, key)
}

func (c *Client) Unlock(key string) (bool, error) {
	return c.Remove(key)
}
