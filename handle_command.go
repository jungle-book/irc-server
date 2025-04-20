package main

import (
	"fmt"
	"log"
)

func handleCommand(client *Client, command string) {
	msg := parseIRCMessage(command)
	log.Println("Message from client: ", client, msg)

	switch msg.Command {
	case "NICK":
		client.nick = msg.Params[0]
		client.ch <- fmt.Sprintf("Welcome, %s!", client.nick)
	default:
		client.ch <- "Unknown command"
	}
}