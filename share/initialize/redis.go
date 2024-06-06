package initialize

import (
	"github.com/redis/go-redis/v9"
	"worframe/share/types"
)

func InitRedis(c *types.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     c.Redis.Host,
		Password: c.Redis.Pass,
		DB:       c.Redis.DB,
	})
}
