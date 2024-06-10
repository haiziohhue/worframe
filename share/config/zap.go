package config

type Zap struct {
	Director     string `  yaml:"director"`      // 日志文件夹
	LogInConsole bool   ` yaml:"log-in-console"` // 输出控制台
	MaxSize      int    ` yaml:"max-size"`       // 输出控制台
	MaxAge       int    ` yaml:"max-age"`        // 输出控制台
	MaxBackups   int    ` yaml:"max-backups"`    // 输出控制台
}
