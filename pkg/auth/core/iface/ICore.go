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
	InitRepository() ICore

	GetEngine() *gin.Engine
	GetRawCore() *shareCoreIface.ICore
	GetAuthConf() *config.AuthPackConfig
	GetRepo(name string) (interface{}, error)
	Run()
}
