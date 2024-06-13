package router

import (
	"github.com/gin-gonic/gin"
	"worframe/pkg/auth/controller"
	"worframe/pkg/auth/core/iface"
)

func RegisterRole(r *gin.Engine, core iface.ICore) {
	role := r.Group("/role")
	ctrl := controller.NewRoleController(core)
	{
		role.GET("/:id", ctrl.GetOne)
		role.GET("", ctrl.GetAll)
		role.PUT("/:id", ctrl.Update)
		role.DELETE("/:id", ctrl.Delete)
		role.POST("", ctrl.Create)
	}
}
