package core

import (
	"github.com/gomodule/redigo/redis"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"log"
	"worframe/share/config"
	"worframe/share/core/iface"
	"worframe/share/utils"
)

type ShareApp struct {
	Conf    *config.Config
	DB      *gorm.DB
	Redis   *redis.Pool
	Logger  *zap.Logger
	WorkDir string
	Err     error
	Env     string
}

var Log *zap.SugaredLogger

func NewApp(env string) iface.ICore {
	app := &ShareApp{}
	if env == "" {
		env = "dev"
	}
	app.Env = env
	workDir, err := utils.FindWorkDir()
	if err != nil {
		app.Err = err
		return nil
	}
	app.WorkDir = workDir
	log.Println("work dir", workDir)
	app.Conf = initConfig(env, workDir)
	return app
}
