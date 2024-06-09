package main

import (
	"worframe/pkg/auth/config"
	initialize2 "worframe/pkg/auth/initialize"
	"worframe/share/core"
	"worframe/share/initialize"
)

func main() {
	core.Cfg = initialize.InitConfig("dev")
	config.AuthCfg = initialize2.InitAuthConfig("dev")
	core.DB = initialize.InitGorm(core.Cfg)
	core.Redis = initialize.InitRedis(core.Cfg)
	core.Engine = initialize.InitGin("auth")
	_ = core.Engine.Run()
}
