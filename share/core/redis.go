package core

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

func (a *ShareApp) InitRedis() *ShareApp {
	if a.Conf == nil {
		a.Error = fmt.Errorf("conf is nil, init redis error")
		return a
	}
	a.Redis = &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", a.Conf.Redis.Host)
			if err != nil {
				return nil, err
			}
			if _, err = conn.Do("AUTH", a.Conf.Redis.Pass); err != nil {
				_ = conn.Close()
				return nil, err
			}
			if _, err = conn.Do("SELECT", a.Conf.Redis.DB); err != nil {
				_ = conn.Close()
				return nil, err
			}
			return conn, nil
		},
	}
	return a
}
