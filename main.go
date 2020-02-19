package main

import (
	"flag"
	"log"
)

func main() {
	service := flag.String("service", "", "The service that will be ran, should be one of either client or server")
	port := ":8080"

	flag.Parse()
	log.Println(*service)

	if *service == "server" {
		startServer(port)
	} else if *service == "client" {
		startClient(port)
	} else {
		log.Fatal("Invalid service flag")
	}

}
