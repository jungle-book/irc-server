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
	case "USER":
		client.username = msg.Params[0]
		client.ch <- fmt.Sprintf("Your username is, %s!", client.username)
	case "JOIN":
		joinChannel(client, msg.Params[0])
	case "LISTCHANNELS":
		listChannels(client)
	case "PRIVMSG":
		handlePrivMessage(client, msg)
	default:
		client.ch <- "Unknown command"
	}
}