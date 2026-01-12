package main

import (
	"VyacheslavKuchumov/test-backend/cmd/api"
	"log"
)

func main() {
	server := api.NewAPIServer(":8000", nil)
	err := server.Run()
	if err != nil {
		log.Fatal(err)
	}

}
