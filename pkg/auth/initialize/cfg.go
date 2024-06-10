package initialize

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"worframe/pkg/auth/config"
	"worframe/share/core"
)

func InitAuthConfig(env string) *config.AuthPackConfig {
	configFileName := fmt.Sprintf("%s.config.yaml", env)
	configFilePath := filepath.Join(core.WorkDir, "cfg", configFileName)
	data, err := os.ReadFile(configFilePath)
	if err != nil {
		panic(err)
	}
	c := &config.AuthPackConfig{
		Config: *core.Cfg,
	}
	err = yaml.Unmarshal(data, &c)
	if err != nil {
		panic(err)
	}
	return c
}
