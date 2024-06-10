package initialize

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
	"worframe/share/core"
	"worframe/share/utils"
)

// InitZap 日志组件初始化
func InitZap(c *config.Config) (logger *zap.SugaredLogger) {
	log.Println(core.WorkDir + c.Zap.Director)
	if ok, _ := utils.PathExists(core.WorkDir + c.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", core.WorkDir+c.Zap.Director)
		_ = os.Mkdir(core.WorkDir+c.Zap.Director, os.ModePerm)
	}
	logMode := zapcore.DebugLevel

	zCore := zapcore.NewCore(getEncoder(c), zapcore.NewMultiWriteSyncer(getWriter(c), zapcore.AddSync(os.Stdout)), logMode)
	return zap.New(zCore).Sugar()
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
func getWriter(c *config.Config) zapcore.WriteSyncer {
	stLogFilePath := filepath.Join(core.WorkDir, c.Zap.Director, time.Now().Format(time.DateOnly)+".log")
	lumberSyncer := &lumberjack.Logger{
		Filename:   stLogFilePath,
		MaxBackups: c.Zap.MaxBackups,
		MaxSize:    c.Zap.MaxSize * (1024 * 1024), //1M
		MaxAge:     c.Zap.MaxAge,
		Compress:   false,
	}
	return zapcore.AddSync(lumberSyncer)
}
