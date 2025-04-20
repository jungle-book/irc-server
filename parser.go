package main

import (
	"strings"
)

type IRCMessage struct {
	Prefix string
	Command string
	Params []string
	Trailing string
	IsChannel bool
}

func parseIRCMessage(raw string) IRCMessage {
	var msg IRCMessage

	// Step 1: Extract prefix if present
	if strings.HasPrefix(raw, ":") {
		split := strings.SplitN(raw[1:], " ", 2)
		msg.Prefix = split[0]
		raw = split[1]
	}

	// Step 2: Extract trailing param (after last " :")
	var trailing string
	if idx := strings.Index(raw, " :"); idx != -1 {
		trailing = raw[idx+2:]
		raw = raw[:idx]
	}

	// Step 3: Split command and parameters
	parts := strings.Fields(raw)
	if len(parts) > 0 {
		msg.Command = parts[0]
		if len(parts) > 1 {
			msg.Params = parts[1:]
		}
	}

	msg.Trailing = trailing
	return msg
}