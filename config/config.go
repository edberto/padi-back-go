package config

import (
	"log"

	"github.com/spf13/viper"
)

type IConfig interface {
	GetString(key string) string
	GetInt(key string) int
}

type Config struct {
	*viper.Viper
}

func NewConfig(path string) IConfig {
	viper.SetConfigFile(path)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Failed to read config file")
	}

	c := viper.GetViper()

	return &Config{c}
}

func (c Config) GetString(key string) string {
	return c.Viper.GetString(key)
}

func (c Config) GetInt(key string) int {
	return c.Viper.GetInt(key)
}
