package config

import (
	"github.com/pupi94/madara/tools/redis"
)

var RedisClient *redis.Client

func InitRedisClient() {
	if RedisClient != nil {
		return
	}
	RedisClient = redis.NewClient(Env.RdsHost, Env.RdsPort, &redis.Config{NameSpace: Env.RdsNamespace})
}
