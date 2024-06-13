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
func (app *ShareApp) GetEnv() string {
	return app.Env
}
func (app *ShareApp) SetEnv(env string) {
	app.Env = env
}
func (app *ShareApp) GetWorkDir() string {
	return app.WorkDir
}
func (app *ShareApp) SetWorkDir(dir string) {
	app.WorkDir = dir
}
func (app *ShareApp) GetConf() *config.Config {
	return app.Conf
}
