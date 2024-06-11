package config

import "worframe/share/config"

type AuthPackConfig struct {
	config.Config
	Captcha `yaml:"captcha"`
}

var AuthCfg *AuthPackConfig
