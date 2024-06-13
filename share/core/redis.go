package core

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
	"worframe/share/core/iface"
)

func (app *ShareApp) InitRedis() iface.ICore {
	if app.Conf == nil {
		app.Err = fmt.Errorf("conf is nil, init redis error")
		return app
	}
	app.Redis = &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", app.Conf.Redis.Host)
			if err != nil {
				return nil, err
			}
			if _, err = conn.Do("AUTH", app.Conf.Redis.Pass); err != nil {
				_ = conn.Close()
				return nil, err
			}
			if _, err = conn.Do("SELECT", app.Conf.Redis.DB); err != nil {
				_ = conn.Close()
				return nil, err
			}
			return conn, nil
		},
	}
	return app
}
func (app *ShareApp) GetRedis() *redis.Pool {
	return app.Redis
}
