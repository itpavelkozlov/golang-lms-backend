package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Service Service
}

type Service struct {
	Name   string `yaml:"name"`
	Logger Logger
}

type Logger struct {
	Cores      []string
	Level      string
	Filename   string
	MaxSize    int
	MaxBackups int
	Compress   bool
}

func NewConfig(path string) (*Config, error) {
	viper.SetConfigFile(path)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	conf := new(Config)
	err = viper.Unmarshal(conf)
	return conf, err
}
