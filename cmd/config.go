package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

const config_path string = "./cache"

func setDefaultValues() {
	viper.SetDefault("api_port", 10000)
}

func InitConfig() {
	setDefaultValues()
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(config_path)
	_ = os.Mkdir(config_path, os.ModePerm)
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		fmt.Println(err)
		viper.SafeWriteConfig()
	}
}
