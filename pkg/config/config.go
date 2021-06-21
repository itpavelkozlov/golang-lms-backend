package config

import (
	"github.com/spf13/viper"
)

func NewConfig(path string) error {
	viper.SetConfigFile(path)
	return viper.ReadInConfig()
}
