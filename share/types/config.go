package types

type Config struct {
	Name   string `yaml:"name"`
	Host   string `yaml:"host"`
	Port   int    `yaml:"port"`
	Server struct {
		Mode       string `yaml:"mode"`
		Secret     string `yaml:"secret"`
		BufferTime string `yaml:"bufferTime"`
		ExpireTime string `yaml:"expireTime"`
	} `yaml:"server"`
	Postgres struct {
		Host string `yaml:"host"`
		User string `yaml:"user"`
		Port int    `yaml:"port"`
		Pass string `yaml:"pass"`
		DB   string `yaml:"database"`
	} `yaml:"postgres"`
	Redis struct {
		Host string `yaml:"host"`
		Pass string `yaml:"pass"`
		DB   int    `yaml:"db"`
	} `yaml:"redis"`
}
