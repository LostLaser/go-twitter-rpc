package main

import (
	"bufio"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"time"

	"github.com/LostLaser/go-rpc-example/message"
)

var port string

func startClient(inpPort string) {
	if len(inpPort) > 0 {
		port = inpPort
	} else {
		port = ":8080"
	}

	// continually read input from console
	for {
		switchInput(readInput("Enter command: "))
	}
}

// getConnection will connect to the rpc
func getConnection() *rpc.Client {
	// run the client
	client, err := rpc.DialHTTP("tcp", "localhost"+port)
	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	return client
}

// readInput will read in a line of text from the terminal
func readInput(prompt string) string {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print(prompt)
	scanner.Scan()

	return scanner.Text()
}

// switchInput determines the input command
func switchInput(input string) {
	switch input {
	case "read":
		handleRead()
	case "write":
		handleWrite()
	default:
		log.Println("Invalid option")
	}
}

// handleRead will take care of the read command
func handleRead() {
	msg := message.Message{}
	var reply *[]message.Message
	getConnection().Call("Task.GetMyMessages", msg, &reply)
	log.Println(reply)
}

// handleWrite will take care of write command, loops until 'exit' is typed
func handleWrite() {
	var text string

	for {
		text = readInput(">")
		if text == "exit" {
			break
		} else {
			msg := message.Message{
				Username:  "Jacob",
				Content:   text,
				TimeStamp: time.Now(),
			}
			var reply *message.Message

			getConnection().Call("Task.SendMessage", msg, &reply)
		}
	}
}
