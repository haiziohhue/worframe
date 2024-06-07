package server

import (
	"github.com/gin-gonic/gin"
	"worframe/pkg/auth/server/middleware"
	"worframe/pkg/auth/server/router"
)

func AuthInitServer(r *gin.Engine) {
	middleware.AuthMiddle(r)
	router.AuthRouter(r)
}
