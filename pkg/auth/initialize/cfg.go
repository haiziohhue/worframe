package initialize

import (
	"gopkg.in/yaml.v3"
	"os"
	"worframe/pkg/auth/config"
	"worframe/share/core"
	"worframe/share/utils"
)

func InitAuthConfig(env string) *config.AuthPackConfig {
	filename, err := utils.FindConfigFile(env, 5)
	if err != nil {
		panic(err)
	}

	data, err := os.ReadFile(filename)
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
