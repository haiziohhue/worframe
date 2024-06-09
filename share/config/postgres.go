package config

type Postgres struct {
	Host  string `yaml:"host"`
	User  string `yaml:"user"`
	Port  int    `yaml:"port"`
	Pass  string `yaml:"pass"`
	DB    string `yaml:"database"`
	Debug bool   `yaml:"debug"`
}
