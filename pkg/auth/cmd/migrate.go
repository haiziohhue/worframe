package main

import (
	authCore "worframe/pkg/auth/core"
	initialize2 "worframe/pkg/auth/initialize"
	"worframe/pkg/auth/migrate"
	"worframe/pkg/auth/service"
	"worframe/share/core"
	"worframe/share/initialize"
)

func main() {
	core.Cfg = initialize.InitConfig("dev")
	core.Logger = initialize.InitZap(core.Cfg)
	core.DB = initialize.InitGorm(core.Cfg)
	core.Redis = initialize.InitRedis(core.Cfg)

	m := migrate.NewDBMigrate(core.DB)
	_ = m.DevEnvInit()

	authCore.Casbin = initialize2.InitCasbin(core.Cfg, core.DB, core.Redis)
	cs := service.NewCasbinService(*authCore.Casbin)
	err := cs.SqlUpdateFlow()
	if err != nil {
		panic(err)
	}
	return
}
