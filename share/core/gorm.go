package core

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"worframe/share/core/iface"
)

func (app *ShareApp) InitDb() iface.ICore {
	if app.Conf == nil {
		app.Err = fmt.Errorf("conf is nil, init database error")
		return app
	}
	dsn := fmt.Sprintf("host=%s user=%s port=%d dbname=%s password=%s sslmode=disable",
		app.Conf.Postgres.Host, app.Conf.Postgres.User, app.Conf.Postgres.Port, app.Conf.Postgres.DB, app.Conf.Postgres.Pass)
	log.Println(dsn)
	pgDb, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}))
	if app.Conf.Postgres.Debug {
		pgDb = pgDb.Debug()
	}
	if err != nil {
		app.Err = err
		return app
	}
	app.DB = pgDb
	return app
}
func (app *ShareApp) GetDB() *gorm.DB {
	return app.DB
}
