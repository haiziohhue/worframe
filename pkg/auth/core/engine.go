package core

import (
	"fmt"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"time"
	"worframe/pkg/auth/core/iface"
	"worframe/pkg/auth/server/middleware"
	"worframe/pkg/auth/server/router"
	shareMiddleware "worframe/share/middleware"
)

func (ac *AuthCore) Run() {
	port := fmt.Sprintf(":%d", ac.GetConf().Server.Port)
	err := ac.Engine.Run(port)
	if err != nil {
		ac.SetErr(err)
		panic(err)
	}
}
func (ac *AuthCore) InitEngine() iface.ICore {
	r := gin.New()
	r.Use(ginzap.Ginzap(ac.GetLog(), time.DateTime, true), gin.Recovery())
	r.Use(middleware.Response())
	r.Use(shareMiddleware.JWTMiddleware(ac.GetRawCore()))
	router.RegisterAuth(r, ac)
	router.RegisterDept(r, ac)
	router.RegisterRole(r, ac)
	router.RegisterMenu(r, ac)
	ac.Engine = r
	return ac
}
func (ac *AuthCore) GetEngine() *gin.Engine {
	return ac.Engine
}
