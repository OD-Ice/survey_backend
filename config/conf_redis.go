package config

import "fmt"

type Redis struct {
	Addr     string `yaml:"addr"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
	PoolSize int    `yaml:"pool_size"`
}

func (r Redis) AddrPort() string {
	return fmt.Sprintf("%s:%d", r.Addr, r.Port)
}
