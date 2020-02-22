package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/LostLaser/go-twitter-rpc/controller"
)

func startServer(port string) {
	task := new(controller.Task)
	err := rpc.Register(task)

	if err != nil {
		log.Fatal("Unable to start server: ", err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Unable to listen: ", err)
	}
	log.Printf("RPC listenting on port %s", port)

	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("Error serving: ", err)
	}
}
