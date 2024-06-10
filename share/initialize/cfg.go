package initialize

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"worframe/share/config"
	"worframe/share/core"
	"worframe/share/utils"
)

func InitConfig(env string) *config.Config {
	workdir, err := utils.FindWorkDir()
	if err != nil {
		panic(err)
	}
	core.WorkDir = workdir
	configFileName := fmt.Sprintf("%s.config.yaml", env)
	configFilePath := filepath.Join(core.WorkDir, "cfg", configFileName)

	data, err := os.ReadFile(configFilePath)
	if err != nil {
		panic(err)
	}
	c := &config.Config{}
	err = yaml.Unmarshal(data, &c)
	if err != nil {
		panic(err)
	}
	return c
}
