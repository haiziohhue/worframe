package config

type Jwt struct {
	SignKey     string `yaml:"signKey"`
	BufferTime  string `yaml:"bufferTime"`
	ExpiresTime string `yaml:"expiresTime"`
}
