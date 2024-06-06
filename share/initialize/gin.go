package initialize

import (
	"github.com/gin-gonic/gin"
	rbac "worframe/pkg/rbac/server"
)

func InitGin(model string) *gin.Engine {
	engine := gin.New()
	switch model {
	case "rbac":
		rbac.RBACIntiServer(engine)
	default:
		engine.Use(gin.Logger(), gin.Recovery())
	}
	return engine
}
