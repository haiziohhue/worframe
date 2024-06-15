package router

import (
	"github.com/gin-gonic/gin"
	"worframe/pkg/auth/controller"
	"worframe/pkg/auth/core/iface"
	"worframe/share/middleware"
)

func RegisterAuth(r *gin.Engine, core iface.ICore) *gin.Engine {
	ctrl := controller.NewAuthController(core)
	jwtFunc, err := middleware.NewAuthRouter(core, ctrl.LoginPkg())
	if err != nil {
		panic(err)
	}
	auth := r.Group("/auth_log")
	{
		auth.POST("/register", ctrl.Register)
		auth.POST("", jwtFunc.LoginHandler)
		auth.DELETE("", jwtFunc.LogoutHandler)
		auth.PUT("/", jwtFunc.RefreshHandler)
	}
	return r
}
