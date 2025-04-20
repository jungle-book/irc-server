package main

import (
	"fmt"
	"log"
	"strings"
)

type IRCMessage struct {
	Source   string
	Command  string
	Params   []string
	Trailing string
}

func (m IRCMessage) Log() {
	log.Println("IRC Message:", "source:", m.Source, "command:", m.Command, "params:", m.Params, "trailing:", m.Trailing)
}

func (m IRCMessage) Encode() string {
	return fmt.Sprintf(":%s %s %s :%s", m.Source, m.Command, strings.Join(m.Params, " "), m.Trailing)
}

func parseIRCMessage(message string) IRCMessage {
	var msg IRCMessage
	msg.Params = make([]string, 0)
	/*
		[:{source} ]{command}[ parameters][ :{trailing}]
	*/

	// remove the source from the message if it has one
	if message[0] == ':' {
		parts := strings.SplitN(message, " ", 2)
		msg.Source = parts[0]
		message = parts[1]
	}

	parts := getToken(message)
	msg.Command = parts[0]
	message = parts[1]

	for message != "" {
		if message[0] == ':' {
			msg.Trailing = strings.TrimSpace(message[1:])
			break
		}

		parts = getToken(message)
		msg.Params = append(msg.Params, parts[0])

		message = parts[1]
	}

	return msg
}

func getToken(s string) []string {
	split := strings.SplitN(s, " ", 2)

	if len(split) == 1 {
		return []string{split[0], ""}
	}

	return split
}
