// Package config package config
package config

import (
	"os"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

var (
	once      sync.Once
	singleton *Config
)

type Config struct {
	Database Database
	Server   Server
}

// GetConfig -.
func GetConfig() *Config {
	if singleton != nil {
		return singleton
	}

	once.Do(func() {
		filePath := "./configs/config.yml"
		fileStat, err := os.Stat(filePath)
		if err != nil || fileStat.IsDir() {
			panic(err)
		}

		newConfig := Config{}
		if fileStat.Size() > 0 {
			err := cleanenv.ReadConfig(filePath, &newConfig)
			if err != nil {
				panic(err)
			}
		}

		if err := cleanenv.ReadEnv(&newConfig); err != nil {
			panic(err)
		}

		singleton = &newConfig
	})

	return singleton
}
