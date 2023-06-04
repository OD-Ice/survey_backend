package config

type Jwt struct {
	Secret string `yaml:"secret"`
	Exp    int    `yaml:"exp"`
}
