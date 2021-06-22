package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Service Service
}

type Service struct {
	Name string `yaml:"name"`
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
