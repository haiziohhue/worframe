package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"worframe/pkg/auth/controller"
)

func RegisterRole(r *gin.Engine, logger *zap.Logger, db *gorm.DB) {
	role := r.Group("/role")
	ctrl := controller.NewRoleController(logger, db)
	{
		role.GET("/:id", ctrl.GetOne)
		role.GET("", ctrl.GetAll)
		role.PUT("/:id", ctrl.Update)
		role.DELETE("/:id", ctrl.Delete)
		role.POST("", ctrl.Create)
	}
}
