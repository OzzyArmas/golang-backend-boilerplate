package config

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func getViber() *viper.Viper {
	v := viper.New()
	v.SetConfigFile("config.yaml")
	v.AddConfigPath(".")
	return v
}

func NewConfig() (*Config, error) {
	log.Info("NewConfig")
	v := getViber()
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %s", err)
	}
	var c Config
	if err := v.Unmarshal(&c); err != nil {
		return nil, fmt.Errorf("error unmarshalling config: %s", err)
	}
	return &c, nil

}
