package core

import (
	"github.com/gin-gonic/gin"
	"worframe/pkg/auth/config"
	"worframe/share/core"
)

type AuthCore struct {
	*core.ShareApp
	Engine   *gin.Engine
	AuthConf *config.AuthPackConfig
}

func NewAuthCore(app *core.ShareApp) *AuthCore {
	AuthApp := &AuthCore{ShareApp: app}
	return AuthApp.initAuthConf()
}
