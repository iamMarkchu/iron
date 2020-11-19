package config

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
)

var config = flag.String("config", "", "config file")

func Init() {
	flag.Parse()
	viper.SetConfigFile(*config)
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {                // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
