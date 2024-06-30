package application

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	redisadapter "github.com/casbin/redis-adapter/v3"
	"github.com/gomodule/redigo/redis"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"worframe/pkg/auth/strategy"
	"worframe/share/config"
	"worframe/share/constant"
)

type CasbinCore struct {
	redis         *redisadapter.Adapter
	postgres      *gormadapter.Adapter
	sqlEnforcer   *casbin.Enforcer
	redisEnforcer *casbin.Enforcer
	logger        *zap.SugaredLogger
	redisPool     *redis.Pool
	model         model.Model
	conf          config.Casbin
	modelName     string
	mapping       func() error
	Error         error
}

func NewCasbinCore(conf config.Casbin, logger *zap.Logger, pool *redis.Pool, db *gorm.DB) (core *CasbinCore) {
	redisAdapt, err := redisadapter.NewAdapterWithPool(pool)
	if err != nil {
		core.Error = err
		return core
	}
	postgresAdapt, err := gormadapter.NewAdapterByDB(db)
	CModel, err := modelBind(conf.ModelName)
	if err != nil {
		core.Error = err
		return core
	}
	mapFunc, err := mapBind(db, logger.Sugar(), conf.ModelName)
	if err != nil {
		core.Error = err
		return core
	}
	sqlEnforcer, err := casbin.NewEnforcer(CModel, postgresAdapt)
	if err != nil {
		core.Error = err
		return core
	}
	redisEnforcer, err := casbin.NewEnforcer(CModel, redisAdapt)
	if err != nil {
		core.Error = err
		return core
	}
	core = &CasbinCore{
		redis:         redisAdapt,
		postgres:      postgresAdapt,
		model:         CModel,
		conf:          conf,
		sqlEnforcer:   sqlEnforcer,
		redisEnforcer: redisEnforcer,
		modelName:     conf.ModelName,
		mapping:       mapFunc,
		logger:        logger.Sugar(),
		redisPool:     pool,
	}
	return core
}
func modelBind(name string) (model.Model, error) {
	ms := constant.CasbinModel[name]
	if ms == "" {
		return nil, fmt.Errorf("no found model" + name)
	}
	m, err := model.NewModelFromString(ms)
	if err != nil {
		return nil, err
	}
	return m, nil
}
func mapBind(db *gorm.DB, log *zap.SugaredLogger, name string) (func() error, error) {
	fn := strategy.CasbinMappingStrategyWithSQL[name]
	if fn == nil {
		return nil, fmt.Errorf("no found mapping" + name)
	}
	return fn(db, log), nil
}

type ICasbinService interface {
	UpdateFlow() error
	CheckFlow() (redisOk, postgresOk bool)
	redisUpdateFormPostgres() error
	postgresUpdateFormSql(db *gorm.DB) error
	redisCheck() bool
	postgresCheck() bool
}
type CasbinService struct {
	*CasbinCore
}

func NewCasbinService(app *CasbinCore) CasbinService {
	return CasbinService{
		app,
	}
}
func (c *CasbinService) redisUpdateFormPostgres() error {
	_, err := c.redisPool.Get().Do("DEL", "casbin_rules")
	if err != nil {
		c.logger.Warn(err)
	}
	return c.redis.SavePolicy(c.sqlEnforcer.GetModel())
}
func (c *CasbinService) postgresUpdateFormSql(db *gorm.DB) error {
	var StrategyExec func() error
	if c.mapping == nil {
		Strategy := strategy.CasbinMappingStrategyWithSQL[c.modelName]
		StrategyExec = Strategy(db, c.logger)
	} else {
		StrategyExec = c.mapping
	}
	return StrategyExec()
}
func (c *CasbinService) postgresCheck() bool {
	c.logger.Info("检测Postgres")
	p := c.sqlEnforcer.GetPolicy()
	if p == nil {
		return false
	}
	return true
}
func (c *CasbinService) redisCheck() bool {
	c.logger.Info("检测Redis")
	p := c.redisEnforcer.GetPolicy()
	if p == nil {
		return false
	}
	return true
}
func (c *CasbinService) UpdateFlow(db *gorm.DB) error {
	redisOk, postgresOk := c.CheckFlow()
	if !postgresOk {
		c.logger.Warn(constant.CASBINPOSTGRESNOREADY)
		err := c.postgresUpdateFormSql(db)
		if err != nil {
			c.logger.Error(constant.CASBINPOSTGRESUPDATEFIAL)
		}
	}
	if !redisOk {
		c.logger.Warn(constant.CASBINREDISNOREADY)
		err := c.redisUpdateFormPostgres()
		if err != nil {
			c.logger.Error(constant.CASBINREDISUPDATEFAIL)
			//todo  后续会做警报措施
		}
	}
	return nil
}
func (c *CasbinService) CheckFlow() (redisOk, postgresOk bool) {
	c.logger.Info("开始检查流程")
	return c.redisCheck(), c.postgresCheck()
}
func (c *CasbinService) SqlUpdateFlow(db *gorm.DB) error {
	if err := c.postgresUpdateFormSql(db); err != nil {
		c.logger.Error(constant.CASBINPOSTGRESUPDATEFIAL, err)
		return err
	}
	if err := c.redisUpdateFormPostgres(); err != nil {
		c.logger.Error(constant.CASBINREDISUPDATEFAIL, err)
		return err
	}
	return nil
}
