package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var instance *viper.Viper

func Init()  {
	instance = viper.New()
	instance.SetConfigName("config")
	instance.SetConfigType("ini")
	instance.AddConfigPath("config")
	err := instance.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func GetInstance() *viper.Viper {
	return instance
}