package router

import (
	"github.com/gin-gonic/gin"
	"worframe/pkg/auth/controller"
	"worframe/pkg/auth/core/iface"
)

func RegisterDept(r *gin.Engine, core iface.ICore) *gin.Engine {
	ctrl := controller.NewDeptController(core)
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
