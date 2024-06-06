package server

import (
	"github.com/gin-gonic/gin"
	"worframe/pkg/rbac/server/middleware"
	"worframe/pkg/rbac/server/router"
)

func RBACIntiServer(r *gin.Engine) {
	middleware.RbacMiddleware(r)
	router.RbacRouter(r)
}
