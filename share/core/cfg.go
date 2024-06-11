package core

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"worframe/share/config"
)

func initConfig(env, dir string) *config.Config {
	configFileName := fmt.Sprintf("%s.config.yaml", env)
	configFilePath := filepath.Join(dir, "cfg", configFileName)

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
