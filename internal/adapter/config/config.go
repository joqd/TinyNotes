package config

import (
	"fmt"
	"sync"

	"github.com/spf13/viper"
)

type (
	Config struct {
		App App
	}

	App struct {
		Debug bool
		Port  int
	}
)

var (
	configInstance Config
	once           sync.Once
)

func Load() *Config {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")

		if err := viper.ReadInConfig(); err != nil {
			panic(fmt.Sprintf("failed to read config: %v", err))
		}

		if err := viper.Unmarshal(&configInstance); err != nil {
			panic(fmt.Sprintf("failed to unmarshal config: %v", err))
		}
	})

	return &configInstance
}
