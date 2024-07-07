package main

import (
	"log"
	"rest-api/internal/app/apiserver"
)

func main() {
	server := apiserver.New()
	if err := server.Start(); err != nil {
		log.Fatal(err)
	} 
}
