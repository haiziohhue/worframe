package core

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"os"
	"path/filepath"
	"time"
	"worframe/share/config"
	"worframe/share/utils"
)

func (a *ShareApp) InitZap() *ShareApp {
	if a.Conf == nil {
		a.Error = fmt.Errorf("conf is nil, init logger error")
		return a
	}
	if ok, _ := utils.PathExists(a.WorkDir + a.Conf.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", a.WorkDir+a.Conf.Zap.Director)
		_ = os.Mkdir(a.WorkDir+a.Conf.Zap.Director, os.ModePerm)
	}
	logMode := zapcore.DebugLevel
	var writer zapcore.WriteSyncer
	if a.Conf.Zap.LogInConsole {
		writer = zapcore.NewMultiWriteSyncer(getWriter(a.Conf, a.WorkDir), zapcore.AddSync(os.Stdout))
	} else {
		writer = getWriter(a.Conf, a.WorkDir)
	}

	zCore := zapcore.NewCore(getEncoder(a.Conf), writer, logMode)
	a.Logger = zap.New(zCore)
	a.SLogger = a.Logger.Sugar()
	return a
}
func (a *ShareApp) InitPublicZap() *ShareApp {
	a.InitZap()
	if a.Error != nil {
		return a
	}
	if a.Logger == nil {
		a.Error = fmt.Errorf("logger is nil, init logger error")
		return a
	}
	Log = a.Logger.Sugar()
	return a
}

func getEncoder(c *config.Config) zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time" //时间记录方式
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Local().Format(time.DateTime))
	}
	return zapcore.NewJSONEncoder(encoderConfig)
}
func getWriter(c *config.Config, dir string) zapcore.WriteSyncer {
	stLogFilePath := filepath.Join(dir, c.Zap.Director, time.Now().Format(time.DateOnly)+".log")
	log.Println("日志路径:", stLogFilePath)
	lumberSyncer := &lumberjack.Logger{
		Filename:   stLogFilePath,
		MaxBackups: c.Zap.MaxBackups,
		MaxSize:    c.Zap.MaxSize * (1024 * 1024), //1M
		MaxAge:     c.Zap.MaxAge,
		Compress:   false,
	}
	return zapcore.AddSync(lumberSyncer)
}
