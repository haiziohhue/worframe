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

	m := migrate.NewDBMigrate(app.DB)

	cs := service.NewCasbinService(service.NewCasbinCore(app.Conf.Casbin, app.Logger, app.Redis, app.DB))
	err := cs.SqlUpdateFlow(app.DB)
	if err != nil {
		panic(err)
	}
	_ = m.DevEnvInit(shareApp.Logger)
	return
}
