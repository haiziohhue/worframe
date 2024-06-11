package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"worframe/pkg/auth/controller"
)

func RegisterMenu(r *gin.Engine, logger *zap.Logger, db *gorm.DB) {
	menu := r.Group("/menu")
	ctrl := controller.NewMenuController(logger, db)
	{
		menu.GET("/:id", ctrl.GetOne)
		menu.GET("", ctrl.GetAll)
		menu.PUT("/:id", ctrl.Update)
		menu.DELETE("/:id", ctrl.Delete)
		menu.POST("", ctrl.Create)
	}
}
