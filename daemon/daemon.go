package daemon

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

func handleGetGameVersions(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func InitDaemon() {
	http.HandleFunc("/game/versions", handleGetGameVersions)
	log.Fatal(
		http.ListenAndServe(
			fmt.Sprintf(":%d", viper.GetInt("api_port")),
			nil,
		))
}
