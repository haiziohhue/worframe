package config

import "worframe/share/config"

type AuthPackConfig struct {
	config.Config
	Jwt     `yaml:"jwt"`
	Captcha `yaml:"captcha"`
}

var AuthCfg *AuthPackConfig
