package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/sknji/alert-api/internal/persist/database"
	"github.com/spf13/viper"
)

type Configuration struct {
	Port     string          `json:"port"`
	Database database.Config `json:"database"`
}

// LoadConfigs loads the application configurations from a local file
func LoadConfigs(configFile string) (conf Configuration, err error) {
	log.Infoln("Loading config file", configFile)
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
