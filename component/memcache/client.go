package memcache

import (
	"fmt"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
)

const (
	// DefaultTimeout is the default socket read/write timeout.
	DefaultTimeout = time.Second

	// DefaultMaxIdleConns is the default maximum number of idle connections
	// kept for any single address.
	DefaultMaxIdleConn = 2
)

type Client struct {
	namespace   string
	defaultTtl  time.Duration
	maxIdleConn int
	client      *memcache.Client
}

func NewClient(servers []string) *Client {
	client := memcache.New(servers...)

	return &Client{
		client:      client,
		defaultTtl:  DefaultTimeout,
		maxIdleConn: DefaultMaxIdleConn,
	}
}

func (c *Client) SetDefaultTtl(t time.Duration) {
	c.defaultTtl = t
}

func (c *Client) SetNamespace(namespace string) {
	c.namespace = namespace
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

func (c *Client) Set(key string, value []byte) error {
	item := &memcache.Item{
		Key:        c.namespaceKey(key),
		Value:      value,
		Expiration: int32(c.defaultTtl.Seconds()),
	}
	if err := c.client.Set(item); err != nil {
		return err
	}
	return nil
}

func (c *Client) SetWitTtl(key string, value []byte, ttl time.Duration) error {
	if ttl <= 0 {
		ttl = c.defaultTtl
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

func (c *Client) namespaceKey(key string) string {
	if c.namespace == "" {
		return key
	}
	return fmt.Sprintf("%s:%s", c.namespace, key)
}
