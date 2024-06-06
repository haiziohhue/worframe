package router

import (
	"github.com/gin-gonic/gin"
	"worframe/pkg/rbac/controller"
)

func routerRole(r *gin.Engine) {
	role := r.Group("/role")
	ctrl := controller.RoleController{}
	{
		role.GET("/:id", ctrl.GetOne)
		role.GET("", ctrl.GetAll)
		role.PUT("/:id", ctrl.Update)
		role.DELETE("/:id", ctrl.Delete)
		role.POST("", ctrl.Create)
	}
}
