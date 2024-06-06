package main

import (
	"worframe/common/core"
	"worframe/common/initialize"
)

func main() {
	core.Cfg = initialize.InitConfig("dev")
	core.DB = initialize.InitGorm(core.Cfg)
	core.Redis = initialize.InitRedis(core.Cfg)
	core.Engine = initialize.InitGin()
	_ = core.Engine.Run()
}
