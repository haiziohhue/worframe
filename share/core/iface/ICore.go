package iface

import (
	"github.com/gomodule/redigo/redis"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"worframe/share/config"
)

type ICore interface {
	GetErr() error
	SetErr(err error)
	GetLog() *zap.Logger
	GetSLog() *zap.SugaredLogger
	GetDB() *gorm.DB
	GetRedis() *redis.Pool
	GetEnv() string
	SetEnv(env string)
	GetWorkDir() string
	SetWorkDir(dir string)
	GetConf() *config.Config
	InitRedis() ICore
	InitZap() ICore
	InitPublicZap() ICore
	InitDb() ICore
}
