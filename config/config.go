package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

const CACHE_PATH string = "./cache"

func setDefaultValues() {
	viper.SetDefault("api_port", 10000)
}

func InitConfig() {
	setDefaultValues()
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(CACHE_PATH)
	_ = os.Mkdir(CACHE_PATH, os.ModePerm)
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		fmt.Println(err)
		viper.SafeWriteConfig()
	}
}
