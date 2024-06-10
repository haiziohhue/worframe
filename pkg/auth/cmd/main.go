package main

import (
	"worframe/pkg/auth/config"
	authCore "worframe/pkg/auth/core"
	initialize2 "worframe/pkg/auth/initialize"
	"worframe/share/core"
	"worframe/share/initialize"
)

func main() {
	//全局配置
	core.Cfg = initialize.InitConfig("dev")
	//特定配置
	config.AuthCfg = initialize2.InitAuthConfig("dev")
	//日志
	core.Logger = initialize.InitZap(core.Cfg)
	//数据库
	core.DB = initialize.InitGorm(core.Cfg)
	//Redis
	core.Redis = initialize.InitRedis(core.Cfg)
	//casbin
	authCore.Casbin = initialize2.InitCasbin(core.Cfg, core.DB, core.Redis)
	//服务器
	core.Engine = initialize.InitGin("auth")
	//启动
	_ = core.Engine.Run()
}
