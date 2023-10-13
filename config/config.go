package config

import (
	"log"

	"github.com/spf13/viper"
)

type AppConfig struct {
	Environment string `mapstructure:"ENVIRONMENT"`
	DBAddress   string `mapstructure:"DB_ADDRESS"`
	DBPassword  string `mapstructure:"DB_PASSWORD"`
	DBUsername  string `mapstructure:"DB_USERNAME"`
	DBName      string `mapstructure:"DB_NAME"`
	DBPort      int    `mapstructure:"DB_PORT"`
}

var appConfig *AppConfig

func InitConfig() *AppConfig {

	if appConfig != nil {
		return appConfig
	}

	viper.AddConfigPath(".")
	viper.AddConfigPath("../")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Failed read config %v", err)
	}

	err = viper.Unmarshal(&appConfig)

	if err != nil {
		log.Fatalf("Failed unmarshal config %v", err)
	}

	log.Println("Success connected to DB")
	return appConfig
}
