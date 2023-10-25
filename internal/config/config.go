package config

import (
	"github.com/sknji/alert-api/internal/persist/database"
	"github.com/spf13/viper"
	"log"
)

type Configuration struct {
	Port     string          `json:"port"`
	Database database.Config `json:"database"`
}

func LoadConfigs(configFile string) (conf Configuration, err error) {
	log.Println("Loading config file", configFile)
	viper.SetConfigName(configFile)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&conf)
	return
}
