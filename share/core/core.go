package core

import (
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"gorm.io/gorm"
	"worframe/share/config"
)

var (
	DB     *gorm.DB
	Cfg    *config.Config
	Redis  *redis.Pool
	Engine *gin.Engine
)
