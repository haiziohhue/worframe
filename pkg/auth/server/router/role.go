package router

import (
	"github.com/gin-gonic/gin"
	"worframe/pkg/auth/controller"
)

func routerRole(r *gin.Engine) {
	role := r.Group("/role")
	ctrl := controller.NewRoleController()
	{
		role.GET("/:id", ctrl.GetOne)
		role.GET("", ctrl.GetAll)
		role.PUT("/:id", ctrl.Update)
		role.DELETE("/:id", ctrl.Delete)
		role.POST("", ctrl.Create)
	}
}
