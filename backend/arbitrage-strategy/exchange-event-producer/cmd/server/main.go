package main

import (
	"algotrading/arbitrage-strategy/event-producer/internal/config"
	"log"
)
func main() {
	
	_, err := config.New()
	if err != nil {
		log.Fatal(err.Error())	
	}
}