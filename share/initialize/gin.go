package initialize

import (
	"github.com/gin-gonic/gin"
	auth "worframe/pkg/auth/server"
)

func InitGin(model string) *gin.Engine {
	engine := gin.New()
	switch model {
	case "auth":
		auth.AuthInitServer(engine)
	case "mailer":
	case "files":
	default:
		engine.Use(gin.Logger(), gin.Recovery())
	}
	return engine
}
