package main

import (
	"worframe/pkg/auth/migrate"
	"worframe/share/core"
	"worframe/share/initialize"
)

func main() {
	core.Cfg = initialize.InitConfig("dev")
	core.DB = initialize.InitGorm(core.Cfg)
	m := migrate.NewDBMigrate(core.DB.Debug())
	_ = m.DevEnvInit()
}
