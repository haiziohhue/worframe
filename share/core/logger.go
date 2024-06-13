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
	"worframe/share/core/iface"
	"worframe/share/utils"
)

func (app *ShareApp) InitZap() iface.ICore {
	if app.Conf == nil {
		app.SetErr(fmt.Errorf("conf is nil, init logger error"))
		return app
	}
	if ok, _ := utils.PathExists(app.WorkDir + app.Conf.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", app.WorkDir+app.Conf.Zap.Director)
		_ = os.Mkdir(app.WorkDir+app.Conf.Zap.Director, os.ModePerm)
	}
	logMode := zapcore.DebugLevel
	var writer zapcore.WriteSyncer
	if app.Conf.Zap.LogInConsole {
		writer = zapcore.NewMultiWriteSyncer(getWriter(app.Conf, app.WorkDir), zapcore.AddSync(os.Stdout))
	} else {
		writer = getWriter(app.Conf, app.WorkDir)
	}

	zCore := zapcore.NewCore(getEncoder(app.Conf), writer, logMode)
	app.Logger = zap.New(zCore)
	return app
}
func (app *ShareApp) InitPublicZap() iface.ICore {
	app.InitZap()
	if app.GetErr() != nil {
		return app
	}
	if app.Logger == nil {
		app.SetErr(fmt.Errorf("logger is nil, init logger error"))
		return app
	}
	Log = app.Logger.Sugar()
	return app
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
func (app *ShareApp) GetLog() *zap.Logger {
	return app.Logger
}
func (app *ShareApp) GetSLog() *zap.SugaredLogger {
	return app.Logger.Sugar()
}
