package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	InitConfig()
	fmt.Println(viper.GetInt("api_port"))
}
