package middleware

import (
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"time"
	"worframe/share/core"
)

func AuthMiddle(r *gin.Engine) *gin.Engine {
	r.Use(ginzap.Ginzap(core.Logger.Desugar(), time.DateTime, true), gin.Recovery())
	r.Use(Response())
	return r
}
