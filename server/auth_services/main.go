package main

import (
	"fmt"
	"log"
	"music-app/authentication-services/internal/api"
)

func main() {

	server := api.NewServer()

	err := server.Run()
	if err != nil {
		log.Fatal("Failed to run authentication Server: ", err)
	}

	fmt.Println("Authentication Server started successfully ")
}
