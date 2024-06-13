package server

import (
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"time"
	"worframe/pkg/auth/core/iface"
	"worframe/pkg/auth/server/middleware"
	"worframe/pkg/auth/server/router"
)

func InitEngine(core iface.ICore) *gin.Engine {
	r := gin.New()

	r.Use(ginzap.Ginzap(core.GetLog(), time.DateTime, true), gin.Recovery())
	r.Use(middleware.Response())

	router.RegisterDept(r, core)
	router.RegisterRole(r, core)
	router.RegisterMenu(r, core)
	return r
}
