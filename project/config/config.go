package config

import "github.com/spf13/viper"

func NewConfig(name, typeFile, path string) *viper.Viper {
	config := viper.New()
	config.SetConfigName(name)
	config.SetConfigType("json")
	config.AddConfigPath(path)

	return config
}
