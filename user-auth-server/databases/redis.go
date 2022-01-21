package databases

import (
	"github.com/go-redis/redis/v8"

	config "user-auth-server/config"
)

var (
	connRedis *redis.Client
)

func InitRedis() {
	c := config.GetConfig()
	cfg := redis.Options{
		Addr:     c.Get("databases.redis.domain").(string) + ":" + c.Get("databases.redis.port").(string),
		Password: c.Get("databases.redis.password").(string),
		DB:       c.Get("databases.redis.databaseName").(int),
	}

	connRedis = redis.NewClient(&cfg)
}

func GetRedis() *redis.Client {
	return connRedis
}
