package config

import (
	"go.uber.org/zap/zapcore"
	"time"
)

type Zap struct {
	Level         string ` yaml:"level"`          // 级别
	Prefix        string ` yaml:"prefix"`         // 日志前缀
	Format        string ` yaml:"format"`         // 输出
	Director      string `  yaml:"director"`      // 日志文件夹
	EncodeLevel   string ` yaml:"encode-level"`   // 编码级
	StacktraceKey string ` yaml:"stacktrace-key"` // 栈名
	ShowLine      bool   ` yaml:"show-line"`      // 显示行
	LogInConsole  bool   ` yaml:"log-in-console"` // 输出控制台
	MaxSize       int    ` yaml:"max-size"`       // 输出控制台
	MaxAge        int    ` yaml:"max-age"`        // 输出控制台
	MaxBackups    int    ` yaml:"max-backups"`    // 输出控制台
}

func (c *Zap) Levels() []zapcore.Level {
	levels := make([]zapcore.Level, 0, 7)
	level, err := zapcore.ParseLevel(c.Level)
	if err != nil {
		level = zapcore.DebugLevel
	}
	for ; level <= zapcore.FatalLevel; level++ {
		levels = append(levels, level)
	}
	return levels
}
func (c *Zap) Encoder() zapcore.Encoder {
	config := zapcore.EncoderConfig{
		TimeKey:       "time",
		NameKey:       "name",
		LevelKey:      "level",
		CallerKey:     "caller",
		MessageKey:    "message",
		StacktraceKey: c.StacktraceKey,
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeTime: func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
			encoder.AppendString(c.Prefix + t.Format("2006-01-02 15:04:05"))
		},
		EncodeLevel:    c.LevelEncoder(),
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
	}
	if c.Format == "json" {
		return zapcore.NewJSONEncoder(config)
	}
	return zapcore.NewConsoleEncoder(config)

}
func (c *Zap) LevelEncoder() zapcore.LevelEncoder {
	switch {
	case c.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		return zapcore.LowercaseLevelEncoder
	case c.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		return zapcore.LowercaseColorLevelEncoder
	case c.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
		return zapcore.CapitalLevelEncoder
	case c.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		return zapcore.CapitalColorLevelEncoder
	default:
		return zapcore.LowercaseLevelEncoder
	}
}
