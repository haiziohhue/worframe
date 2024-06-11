package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"worframe/pkg/auth/controller"
)

func RegisterDept(r *gin.Engine, logger *zap.Logger, db *gorm.DB) *gin.Engine {
	ctrl := controller.NewDeptController(logger, db)
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
