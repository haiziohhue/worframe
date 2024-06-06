package router

import (
	"github.com/gin-gonic/gin"
	"worframe/pkg/rbac/controller"
)

func routerDept(r *gin.Engine) *gin.Engine {
	ctrl := controller.DeptController{}
	dept := r.Group("/dept")
	{
		dept.GET("", ctrl.GetAll)
		dept.POST("", ctrl.Create)
		dept.GET("/:id", ctrl.GetOne)
		dept.PUT("/:id", ctrl.Update)
		dept.DELETE("/:id", ctrl.Delete)
	}
	return r
}
