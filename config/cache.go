package config

import (
	"fmt"
	"github.com/pupi94/madara/tools/memcache"
	"github.com/pupi94/madara/tools/redis"
	"time"
)

var (
	RedisClient    *redis.Client
	MemcacheClient *memcache.Client
)

func InitRedisClient() {
	if RedisClient != nil {
		return
	}
	RedisClient = redis.NewClient(fmt.Sprintf("%s:%s", Env.RdsHost, Env.RdsPort))
	RedisClient.SetDefaultTtl(24 * time.Hour * 3)
	RedisClient.SetNamespace(Env.RdsNamespace)
}

func InitMemcacheClient() {
	if MemcacheClient != nil {
		return
	}
	servers := []string{
		fmt.Sprintf("%s:%s", Env.RdsHost, Env.RdsPort),
	}
	MemcacheClient = memcache.NewClient(servers)
	MemcacheClient.SetDefaultTtl(24 * time.Hour * 3)
	MemcacheClient.SetNamespace(Env.RdsNamespace)
}
