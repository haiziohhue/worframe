package initialize

import (
	"gopkg.in/yaml.v3"
	"os"
	"worframe/share/config"
	"worframe/share/utils"
)

func InitConfig(env string) *config.Config {
	filename, err := utils.FindConfigFile(env, 5)
	data, err := os.ReadFile(filename)
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
