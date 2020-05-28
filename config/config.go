package config

import "github.com/spf13/viper"

type IConfig interface {
	GetString(key string) string
}

type Config struct {
	*viper.Viper
}

func NewConfig(path string) IConfig {
	viper.SetConfigFile(path)

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	c := viper.GetViper()

	return &Config{c}
}

func (c Config) GetString(key string) string {
	return c.Viper.GetString(key)
}
