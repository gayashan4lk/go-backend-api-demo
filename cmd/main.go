package main

import (
	"log"

	"github.com/gayashan4lk/go-backend-api-demo/cmd/api"
)

func main() {
	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
