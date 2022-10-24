package config

import (
	"fmt"
	"github.com/spf13/viper"
	"go-chat/pkg/dbclient/postgres"
)

type Config struct {
	DB postgres.Config `mapstructure:"db"`
}

func Load() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("can't config: %v", err)
	}

	var config Config

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("can't unmarshal config: %v", err)
	}

	return &config, nil
}
