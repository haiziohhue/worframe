package config

type Redis struct {
	Host string `yaml:"host"`
	Pass string `yaml:"pass"`
	DB   int    `yaml:"db"`
}
