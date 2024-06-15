package core

import (
	"github.com/gin-gonic/gin"
	"worframe/pkg/auth/config"
	"worframe/pkg/auth/core/iface"
	shareCoreIface "worframe/share/core/iface"
)

type AuthCore struct {
	shareCoreIface.ICore
	Engine   *gin.Engine
	AuthConf *config.AuthPackConfig
	Repo     map[string]interface{}
}

func NewAuthCore(app shareCoreIface.ICore) iface.ICore {
	AuthApp := &AuthCore{ICore: app}
	return AuthApp.InitAuthConf()
}
func (ac *AuthCore) GetRawCore() *shareCoreIface.ICore {
	return &ac.ICore
}
