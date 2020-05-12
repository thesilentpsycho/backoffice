package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type LoggingConfig struct {
	Default LoggerConfig   `mapstructure:"default"`
	Loggers []LoggerConfig `mapstructure:"loggers"`
}

type LoggerConfig struct {
	Name  string `mapstructure:"name"`
	File  string `mapstructure:"file"`
	Level string `mapstructure:"level"`
}

type DataBaseConfig struct {
	Host   string `mapstructure:"host"`
	Driver string `mapstructure:"driver"`
}

type GeneralConfig struct {
	Env      string         `mapstructure:"env"`
	Database DataBaseConfig `mapstructure:"database"`
	Logging  LoggingConfig  `mapstructure:"logging"`
}

func Loadconfig(filepath string) *GeneralConfig {

	config := &GeneralConfig{}

	viper.SetConfigFile(filepath)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Failed to read config file: %s \n", err))
	}
	err = viper.Unmarshal(config)
	if err != nil {
		panic(fmt.Errorf("Failed to deserialize config file: %s \n", err))
	}

	return config
}
