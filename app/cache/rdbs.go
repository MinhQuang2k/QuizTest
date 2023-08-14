package cache

import (
	"quiztest/app/interfaces"
	"quiztest/pkg/redis"

	"quiztest/config"
)

type cache struct {
	rdbs redis.IRedis
}

func NewCache() interfaces.ICache {
	cfg := config.GetConfig()
	rdb := redis.New(redis.Config{
		Address:  cfg.RedisURI,
		Password: cfg.RedisPassword,
		Database: cfg.RedisDB,
	})

	return &cache{
		rdbs: rdb,
	}
}

// GetInstance get database instance
func (c *cache) GetInstance() redis.IRedis {
	return c.rdbs
}
