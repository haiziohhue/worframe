package migrate

import (
	"gorm.io/gorm"
	"worframe/share/core"
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
	core.Logger.Info("开始迁移...")
	core.Logger.Info("初始化执行命令")
	err := m.initTend()
	if err != nil {
		return err
	}
	core.Logger.Info("初始化表单")
	err = m.initTable()
	if err != nil {
		return err
	}
	core.Logger.Info("初始化初始化表单数据")
	err = m.initDevData()
	return nil
}
