package service

import (
	"fmt"
	"gorm.io/gorm"
	authCore "worframe/pkg/auth/core"
	"worframe/pkg/auth/strategy"
	"worframe/share/constant"
	"worframe/share/core"
)

type ICasbinService interface {
	UpdateFlow() error
	CheckFlow() (redisOk, postgresOk bool)
	redisUpdateFormPostgres() error
	postgresUpdateFormSql(db *gorm.DB) error
	redisCheck() bool
	postgresCheck() bool
}
type CasbinService struct {
	*authCore.CasbinCore
}

func NewCasbinService(casbinCore authCore.CasbinCore) CasbinService {
	return CasbinService{
		authCore.Casbin,
	}
}
func (c *CasbinService) redisUpdateFormPostgres() error {
	_, err := core.Redis.Get().Do("DEL", "casbin_rules")
	if err != nil {
		core.Logger.Warn(err)
	}
	return c.Redis.SavePolicy(c.SqlEnforcer.GetModel())
}
func (c *CasbinService) postgresUpdateFormSql(db *gorm.DB) error {
	Strategy := strategy.CasbinMappingStrategyWithSQL[c.ModelName]
	if Strategy == nil {
		return fmt.Errorf(constant.CASBINSTRATEGYNOFOUND + c.ModelName)
	}
	return Strategy(db)
}
func (c *CasbinService) postgresCheck() bool {
	core.Logger.Info("检测Postgres")
	p := c.SqlEnforcer.GetPolicy()
	if p == nil {
		return false
	}
	return true
}
func (c *CasbinService) redisCheck() bool {
	core.Logger.Info("检测Redis")
	p := c.RedisEnforcer.GetPolicy()
	if p == nil {
		return false
	}
	return true
}
func (c *CasbinService) UpdateFlow() error {
	redisOk, postgresOk := c.CheckFlow()
	if !postgresOk {
		core.Logger.Warn(constant.CASBINPOSTGRESNOREADY)
		err := c.postgresUpdateFormSql(core.DB)
		if err != nil {
			core.Logger.Error(constant.CASBINPOSTGRESUPDATEFIAL)
		}
	}
	if !redisOk {
		core.Logger.Warn(constant.CASBINREDISNOREADY)
		err := c.redisUpdateFormPostgres()
		if err != nil {
			core.Logger.Error(constant.CASBINREDISUPDATEFAIL)
			//todo  后续会做警报措施
		}
	}

	return nil
}
func (c *CasbinService) CheckFlow() (redisOk, postgresOk bool) {
	core.Logger.Info("开始检查流程")
	return c.redisCheck(), c.postgresCheck()
}
func (c *CasbinService) SqlUpdateFlow() error {
	if err := c.postgresUpdateFormSql(core.DB); err != nil {
		core.Logger.Error(constant.CASBINPOSTGRESUPDATEFIAL)
		return err
	}
	if err := c.redisUpdateFormPostgres(); err != nil {
		core.Logger.Error(constant.CASBINREDISUPDATEFAIL)
		return err
	}
	return nil
}
