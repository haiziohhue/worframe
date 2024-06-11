package server

import (
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
	"worframe/pkg/auth/server/middleware"
	"worframe/pkg/auth/server/router"
)

func InitEngine(logger *zap.Logger, db *gorm.DB) *gin.Engine {
	r := gin.New()

	r.Use(ginzap.Ginzap(logger, time.DateTime, true), gin.Recovery())
	r.Use(middleware.Response())

	router.RegisterDept(r, logger, db)
	router.RegisterRole(r, logger, db)
	router.RegisterMenu(r, logger, db)
	return r
}
