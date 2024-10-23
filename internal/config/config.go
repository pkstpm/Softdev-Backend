package config

import (
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

type (
	Config struct {
		Server   *Server
		Database *Database
		Jwt      *Jwt
	}

	Server struct {
		Prefix string `mapstructure:"prefix" validate:"required"`
		Port   int    `mapstructure:"port" validate:"required"`
	}

	Database struct {
		Host     string `mapstructure:"host" validate:"required"`
		Port     int    `mapstructure:"port" validate:"required"`
		User     string `mapstructure:"user" validate:"required"`
		Password string `mapstructure:"password" validate:"required"`
		Name     string `mapstructure:"name" validate:"required"`
		SSLMode  string `mapstructure:"sslmode" validate:"required"`
		TimeZone string `mapstructure:"timezone" validate:"required"`
	}

	Jwt struct {
		AccessSecret string `mapstructure:"access_secret" validate:"required"`
	}
)

var (
	once           sync.Once
	configInstance *Config
)

// GetConfig initializes and returns the configuration
func GetConfig() *Config {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./cmd/app/config")
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		// Attempt to read the config file
		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Error reading config file: %v", err)
		}

		fmt.Println("Using config file:", viper.ConfigFileUsed())

		// Unmarshal the config into configInstance
		if err := viper.Unmarshal(&configInstance); err != nil {
			log.Fatalf("Unable to unmarshal config: %v", err)
		}
	})

	return configInstance
}
