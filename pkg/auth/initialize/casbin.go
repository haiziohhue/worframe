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
	/*
		1. redis get看看有没有casbin_rules已经存在
		2. 如果redis不存在casbin_rules,从postgres数据库拉取
		3. 如果
	*/
	redisAdapt, err := redisadapter.NewAdapterWithPool(redisPool)
	if err != nil {
		panic(err)
	}
	postgresAdapt, err := gormadapter.NewAdapterByDB(db)
	model := modelBind(c.Casbin.ModelName)
	sqlEnforcer, err := casbin.NewEnforcer(model, postgresAdapt)
	if err != nil {
		panic(err)
	}
	redisEnforcer, err := casbin.NewEnforcer(model, redisAdapt)
	if err != nil {
		panic(err)
	}
	casbin.NewEnforcer(model, redisAdapt)
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
