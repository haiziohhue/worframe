package core

import (
	"github.com/gomodule/redigo/redis"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"log"
	"worframe/share/config"
	"worframe/share/utils"
)

type ShareApp struct {
	Conf    *config.Config
	DB      *gorm.DB
	Redis   *redis.Pool
	Logger  *zap.Logger
	SLogger *zap.SugaredLogger
	WorkDir string
	Error   error
	Env     string
}

var Log *zap.SugaredLogger

func NewApp(env string) (app *ShareApp) {
	if env == "" {
		env = "dev"
	}
	str, err := utils.FindWorkDir()
	if err != nil {
		return &ShareApp{
			Error: err,
		}
	}
	log.Println("work dir", str)
	return &ShareApp{
		Conf:    initConfig(env, str),
		Env:     env,
		WorkDir: str,
	}
}
