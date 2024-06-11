package core

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func (a *ShareApp) InitDb() *ShareApp {
	if a.Conf == nil {
		a.Error = fmt.Errorf("conf is nil, init database error")
		return a
	}
	dsn := fmt.Sprintf("host=%s user=%s port=%d dbname=%s password=%s sslmode=disable",
		a.Conf.Postgres.Host, a.Conf.Postgres.User, a.Conf.Postgres.Port, a.Conf.Postgres.DB, a.Conf.Postgres.Pass)
	log.Println(dsn)
	pgDb, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}))
	if a.Conf.Postgres.Debug {
		pgDb = pgDb.Debug()
	}
	if err != nil {
		a.Error = err
		return a
	}
	a.DB = pgDb
	return a
}
