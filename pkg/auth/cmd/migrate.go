package main

import (
	"worframe/pkg/auth/application"
	authCore "worframe/pkg/auth/core"
	"worframe/pkg/auth/migrate"
	"worframe/share/core"
)

func main() {

	shareApp := core.
		NewApp("test").
		InitPublicZap().
		InitDb().
		InitRedis()

	app := authCore.
		NewAuthCore(shareApp)

	m := migrate.NewDBMigrate(app.GetDB())

	cs := application.NewCasbinService(application.NewCasbinCore(app.GetConf().Casbin, app.GetLog(), app.GetRedis(), app.GetDB()))
	err := cs.SqlUpdateFlow(app.GetDB())
	if err != nil {
		panic(err)
	}
	_ = m.DevEnvInit(shareApp.GetLog())
	return
}
