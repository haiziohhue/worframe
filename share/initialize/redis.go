package initialize

import (
	"github.com/gomodule/redigo/redis"
	"time"
	"worframe/share/config"
)

func InitRedis(c *config.Config) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", c.Redis.Host)
			if err != nil {
				return nil, err
			}
			if _, err = conn.Do("AUTH", c.Redis.Pass); err != nil {
				_ = conn.Close()
				return nil, err
			}
			if _, err = conn.Do("SELECT", c.Redis.DB); err != nil {
				_ = conn.Close()
				return nil, err
			}
			return conn, nil
		},
	}
}
