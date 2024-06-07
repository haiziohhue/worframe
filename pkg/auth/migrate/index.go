package migrate

import (
	"gorm.io/gorm"
)

type DBMigrate struct {
	db *gorm.DB
}

func NewDBMigrate(db *gorm.DB) *DBMigrate {
	return &DBMigrate{db: db}
}
func (m *DBMigrate) TestEnvInit() error {
	err := m.initTend()
	if err != nil {
		return err
	}
	err = m.initTable()
	if err != nil {
		return err
	}
	err = m.initTestData()
	if err != nil {
		return err
	}
	return nil
}
func (m *DBMigrate) DevEnvInit() error {
	err := m.initTend()
	if err != nil {
		return err
	}
	err = m.initTable()
	if err != nil {
		return err
	}
	return nil
}
