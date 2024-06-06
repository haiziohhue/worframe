package middleware

import "github.com/gin-gonic/gin"

func RbacMiddleware(r *gin.Engine) *gin.Engine {
	r.Use(gin.Logger(), gin.Recovery())
	r.Use(Response())
	return r
}
