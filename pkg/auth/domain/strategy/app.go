package strategy

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var CasbinMappingStrategyWithSQL = map[string]func(db *gorm.DB, Logger *zap.SugaredLogger) func() error{
	"rbacWithPattern": rbacWithPattern,
}
