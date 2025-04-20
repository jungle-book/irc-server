package main

import (
	"fmt"
	"strings"
)

func handleCommand(client *Client, command string) {
	parts := strings.SplitN(command, " ", 2)
	cmd := strings.ToUpper(parts[0])

	arg := ""
	if len(parts) > 1 {
		arg = parts[1]
	}

	switch cmd {
	case "NICK":
		client.nick = arg
		client.ch <- fmt.Sprintf("Welcome, %s!", client.nick)
	default:
		client.ch <- "Unknown command"
	}

}