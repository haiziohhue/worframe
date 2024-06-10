package initialize

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	redisadapter "github.com/casbin/redis-adapter/v3"
	"github.com/gomodule/redigo/redis"
	"gorm.io/gorm"
	"worframe/pkg/auth/core"
	"worframe/share/config"
	"worframe/share/constant"
)

func InitCasbin(c *config.Config, db *gorm.DB, redisPool *redis.Pool) *core.CasbinCore {
	redisAdapt, err := redisadapter.NewAdapterWithPool(redisPool)
	if err != nil {
		panic(err)
	}
	postgresAdapt, err := gormadapter.NewAdapterByDB(db)
	model := modelBind(c.Casbin.ModelName)
	if model == nil {
		panic("model bind error")
	}
	sqlEnforcer, err := casbin.NewEnforcer(model, postgresAdapt)
	if err != nil {
		panic(err)
	}
	redisEnforcer, err := casbin.NewEnforcer(model, redisAdapt)
	if err != nil {
		panic(err)
	}
	return &core.CasbinCore{
		Redis:         redisAdapt,
		Postgres:      postgresAdapt,
		Model:         model,
		SqlEnforcer:   sqlEnforcer,
		RedisEnforcer: redisEnforcer,
		ModelName:     c.Casbin.ModelName,
	}
}
func modelBind(name string) model.Model {
	ms := constant.CasbinModel[name]
	if ms == "" {
		panic("no found model" + name)
	}
	m, err := model.NewModelFromString(ms)
	if err != nil {
		panic(err)
	}
	return m
}
