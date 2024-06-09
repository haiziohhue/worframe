package config

type Server struct {
	Mode string `yaml:"mode"`
	Name string `yaml:"name"`
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}
