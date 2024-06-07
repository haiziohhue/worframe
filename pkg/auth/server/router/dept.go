package router

import (
	"github.com/gin-gonic/gin"
	"worframe/pkg/auth/controller"
)

func routerDept(r *gin.Engine) *gin.Engine {
	ctrl := controller.NewDeptController()
	dept := r.Group("/dept")
	{
		dept.GET("/:id", ctrl.GetOne)
		dept.GET("", ctrl.GetAll)
		dept.POST("", ctrl.Create)
		dept.PUT("/:id", ctrl.Update)
		dept.DELETE("/:id", ctrl.Delete)
	}
	return r
}
