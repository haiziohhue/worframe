package main

import (
	authCore "worframe/pkg/auth/core"
	"worframe/pkg/auth/migrate"
	"worframe/pkg/auth/service"
	"worframe/share/core"
)

func main() {

	shareApp := core.
		NewApp("dev").
		InitPublicZap().
		InitDb().
		InitRedis()

	app := authCore.
		NewAuthCore(shareApp)

	m := migrate.NewDBMigrate(app.GetDB())

	cs := service.NewCasbinService(service.NewCasbinCore(app.GetConf().Casbin, app.GetLog(), app.GetRedis(), app.GetDB()))
	err := cs.SqlUpdateFlow(app.GetDB())
	if err != nil {
		panic(err)
	}
	_ = m.DevEnvInit(shareApp.GetLog())
	return
}
