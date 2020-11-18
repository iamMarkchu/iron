package config

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
)

var instance *viper.Viper

var config = flag.String("config", "", "config file")

func Init() {
	flag.Parse()
	instance = viper.New()
	instance.SetConfigFile(*config)
	err := instance.ReadInConfig() // Find and read the config file
	if err != nil {                // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func GetInstance() *viper.Viper {
	return instance
}
