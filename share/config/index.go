package config

type Config struct {
	Postgres `yaml:"postgres"`
	Redis    `yaml:"redis"`
	Server   `yaml:"server"`
	Casbin   `yaml:"casbin"`
	Zap      `yaml:"zap"`
	Jwt      `yaml:"jwt"`
}
