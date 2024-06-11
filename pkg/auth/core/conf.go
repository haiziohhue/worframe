package core

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"worframe/pkg/auth/config"
)

func (ac *AuthCore) initAuthConf() *AuthCore {
	if ac.Env == "" {
		ac.Env = "dev"
	}
	configFileName := fmt.Sprintf("%s.config.yaml", ac.Env)
	configFilePath := filepath.Join(ac.WorkDir, "cfg", configFileName)
	data, err := os.ReadFile(configFilePath)
	if err != nil {
		panic(err)
	}
	c := &config.AuthPackConfig{
		Config: *ac.Conf,
	}
	err = yaml.Unmarshal(data, &c)
	if err != nil {
		panic(err)
	}
	ac.AuthConf = c
	return ac
}
