package main

import (
	"worframe/share/core"
	"worframe/share/initialize"
)

func main() {
	core.Cfg = initialize.InitConfig("dev")
	core.DB = initialize.InitGorm(core.Cfg)
	core.Redis = initialize.InitRedis(core.Cfg)
	core.Engine = initialize.InitGin("rbac")
	_ = core.Engine.Run()
}
