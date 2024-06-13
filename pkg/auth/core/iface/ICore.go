package iface

import (
	"github.com/gin-gonic/gin"
	"worframe/pkg/auth/config"
	shareCoreIface "worframe/share/core/iface"
)

type ICore interface {
	shareCoreIface.ICore
	InitAuthConf() ICore
	InitEngine() ICore
	Run()
	GetEngine() *gin.Engine
	GetRawCore() *shareCoreIface.ICore
	GetAuthConf() *config.AuthPackConfig
}
