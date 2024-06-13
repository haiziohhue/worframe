package router

import (
	"github.com/gin-gonic/gin"
	"worframe/pkg/auth/controller"
	"worframe/pkg/auth/core/iface"
)

func RegisterMenu(r *gin.Engine, core iface.ICore) {
	menu := r.Group("/menu")
	ctrl := controller.NewMenuController(core)
	{
		menu.GET("/:id", ctrl.GetOne)
		menu.GET("", ctrl.GetAll)
		menu.PUT("/:id", ctrl.Update)
		menu.DELETE("/:id", ctrl.Delete)
		menu.POST("", ctrl.Create)
	}
}
