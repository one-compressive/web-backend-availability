package main

import (
	"log"

	"github.com/one-compressive/web-backend-availability/internal/api"
)

func main() {
	log.Println("Application started!")
	api.StartServer()
	log.Println("Application terminated!")
}
