package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct{
	Server ServerConfig `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
}

type ServerConfig struct{
	Port int `mapstructure:"port"`
	Path string `mapstructure:"path"`
	HeartbeatInterval int `mapstructure:"heartbeat_interval"`
}

type DatabaseConfig struct{
	Uri string `mapstructture:"uri"`
	DatabaseName string `mapstructure:"database_name"`
}

var ConfigApp *Config

func Load() (error){
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("error reading config file: %w", err)
	}
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return fmt.Errorf("unable to decode into struct: %w", err)
	}
	ConfigApp = &config
	return nil
}