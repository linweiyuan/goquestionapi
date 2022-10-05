package util

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	configFileName   = "app"
	configFileSuffix = "env"
)

type Config struct {
	ServerPort string `mapstructure:"SERVER_PORT"`

	DBDriver string `mapstructure:"DB_DRIVER"`
	DBSource string `mapstructure:"DB_SOURCE"`

	LogLevel string `mapstructure:"LOG_LEVEL"`
}

func LoadConfig(path string) (config Config) {
	viper.AddConfigPath(path)
	viper.SetConfigName(configFileName)
	viper.SetConfigType(configFileSuffix)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("failed to load config file: [%v]", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("failed to parse config file: [%v]", err)
	}

	logLevel, err := log.ParseLevel(config.LogLevel)
	if err != nil {
		log.Fatalf("failed to get log level: [%v]", err)
	}

	log.SetLevel(logLevel)

	return
}
