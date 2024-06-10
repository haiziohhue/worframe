package strategy

import "gorm.io/gorm"

var CasbinMappingStrategyWithSQL = map[string]func(db *gorm.DB) error{
	"rbacWithPattern": rbacWithPattern,
}
