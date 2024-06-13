package core

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"worframe/pkg/auth/config"
	"worframe/pkg/auth/core/iface"
)

func (ac *AuthCore) InitAuthConf() iface.ICore {
	if ac.GetEnv() == "" {
		ac.SetEnv("dev")
	}
	configFileName := fmt.Sprintf("%s.config.yaml", ac.GetEnv())
	configFilePath := filepath.Join(ac.GetWorkDir(), "cfg", configFileName)
	data, err := os.ReadFile(configFilePath)
	if err != nil {
		panic(err)
	}
	c := &config.AuthPackConfig{
		Config: *ac.GetConf(),
	}
	err = yaml.Unmarshal(data, &c)
	if err != nil {
		panic(err)
	}
	ac.AuthConf = c
	return ac
}
func (ac *AuthCore) GetAuthConf() *config.AuthPackConfig {
	return ac.AuthConf
}
