package core

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	redisadapter "github.com/casbin/redis-adapter/v3"
)

var Casbin *CasbinCore

type CasbinCore struct {
	Redis         *redisadapter.Adapter
	Postgres      *gormadapter.Adapter
	Model         model.Model
	SqlEnforcer   *casbin.Enforcer
	RedisEnforcer *casbin.Enforcer
	ModelName     string
}
