package middleware

import (
	"github.com/gin-gonic/gin"
	"worframe/share/middware"
)

func AuthMiddle(r *gin.Engine) *gin.Engine {
	r.Use(gin.Logger(), gin.Recovery())
	r.Use(Response())
	r.Use(middware.Casbin())
	return r
}
