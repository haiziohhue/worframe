package router

import (
	"github.com/gin-gonic/gin"
	"worframe/pkg/rbac/controller"
)

func routerMenu(r *gin.Engine) {
	menu := r.Group("/menu")
	ctrl := controller.MenuController{}
	{
		menu.GET("/:id", ctrl.GetOne)
		menu.GET("", ctrl.GetAll)
		menu.PUT("/:id", ctrl.Update)
		menu.DELETE("/:id", ctrl.Delete)
		menu.POST("", ctrl.Create)
	}
}
