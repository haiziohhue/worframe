package initialize

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"worframe/share/types"
)

func InitGorm(c *types.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s port=%d dbname=%s password=%s sslmode=disable",
		c.Postgres.Host, c.Postgres.User, c.Postgres.Port, c.Postgres.DB, c.Postgres.Pass)
	fmt.Println(dsn)
	pgDb, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}))
	if err != nil {
		panic(err)
	}
	return pgDb
}
