package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/fledgexu/mineconsole/daemon"
	"github.com/spf13/viper"
)

func main() {
	InitConfig()
	fmt.Println(viper.GetInt("api_port"))
	go daemon.InitDaemon()
	// wait for a SIGINT or SIGTERM signal
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	fmt.Println("Received signal, shutting down...")
}
