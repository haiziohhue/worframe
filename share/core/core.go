package core

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"worframe/share/types"
)

var (
	DB     *gorm.DB
	Cfg    *types.Config
	Redis  *redis.Client
	Engine *gin.Engine
)
