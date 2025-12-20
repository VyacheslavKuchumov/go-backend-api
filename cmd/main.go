package main

import (
	"log"
	"vyachik/test-backend/cmd/api"
)

func main() {
	server := api.NewAPIServer(":8000", nil)
	err := server.Run()
	if err != nil {
		log.Fatal(err)
	}

}
