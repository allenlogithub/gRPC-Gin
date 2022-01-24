package config

import (
	"fmt"
	"log"

	viper "github.com/spf13/viper"
)

var config *viper.Viper

func Init(env string) {
	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName(env)
	config.AddConfigPath("../config/")
	config.AddConfigPath("config/")
	if err := config.ReadInConfig(); err != nil {
		// add notification, eg, email
		fmt.Println(err)
		log.Fatal("error on parsing configuration file")
	}
}

func GetConfig() *viper.Viper {
	return config
}
