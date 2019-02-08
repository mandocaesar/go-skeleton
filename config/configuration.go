package config

import (
	"github.com/spf13/viper"
)

//Configuration : configuration struct
type Configuration struct {
	Database Database
	Server   Server
	Loggly   Loggly
}

//New : instantiate Configuration
func New(customPath string) (*Configuration, error) {
	path := "./../"
	if customPath != "" {
		path = customPath
	}

	viper.SetConfigName("default")
	viper.AddConfigPath(path)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := new(Configuration)
	if err := viper.Unmarshal(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
