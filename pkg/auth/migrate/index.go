package migrate

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type DBMigrate struct {
	db *gorm.DB
}

func NewDBMigrate(db *gorm.DB) *DBMigrate {
	return &DBMigrate{db: db}
}
func (m *DBMigrate) TestEnvInit(logger *zap.Logger) error {
	logger.Info("开发数据开始迁移...")
	logger.Info("初始化执行命令")
	err := m.initTend()
	if err != nil {
		return err
	}
	logger.Info("初始化表单")
	err = m.initTable()
	if err != nil {
		return err
	}
	logger.Info("初始化初始化表单数据")
	err = m.initTestData()
	if err != nil {
		return err
	}
	return nil
}
func (m *DBMigrate) DevEnvInit(logger *zap.Logger) error {
	logger.Info("开发数据开始迁移...")
	logger.Info("初始化执行命令")
	err := m.initTend()
	if err != nil {
		return err
	}
	logger.Info("初始化表单")
	err = m.initTable()
	if err != nil {
		return err
	}
	logger.Info("初始化初始化表单数据")
	err = m.initDevData()
	return nil
}
