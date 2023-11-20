package main

import (
	"algotrading/common-service/transfer-service/internal/config"
	"algotrading/common-service/transfer-service/internal/health"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	config, err := config.Load()
	if err != nil {
		log.Printf("failed to load application configuration: %s", err)
		os.Exit(-1)
	}

	r := attachHandlers(config)

	err = http.ListenAndServe(fmt.Sprintf(":%d", config.ServerPort), r)
	if err != nil {
		log.Fatalln("there's an error with the server: ", err)
	}
}

func attachHandlers(config *config.Config) *mux.Router {
	r := mux.NewRouter()
	health.RegisterHandlers(r)

	return r
}
