package initialize

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"worframe/share/types"
)

func InitConfig(env string) *types.Config {
	filename := fmt.Sprintf("cfg/%s.config.yaml", env)
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	config := &types.Config{}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}
	return config
}
