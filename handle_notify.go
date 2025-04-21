package main

import (
	"fmt"
	"strings"
)

func handleNotifyCommand(client *Client, msg IRCMessage) {
	target := msg.Params[0]

	mutex.Lock()
	defer mutex.Unlock()

	resp := IRCMessage{
		Source:   client.nick,
		Command:  msg.Command,
		Params:   msg.Params,
		Trailing: msg.Trailing,
	}

	if strings.Contains(msg.Params[0], "#") {
		if members, exists := channels[target]; exists {
			for member := range members {
				if member.nick != client.nick {
					resp.Params = []string{member.nick}
					member.ch <- resp.Encode()
				}
			}
		}
	} else {
		recipent, ok := clients[target]

		if ok {
			recipent.ch <- fmt.Sprintf(":%s PRIVMSG %s %s", client.nick, target, msg.Trailing)
			return
		}

		client.ch <- fmt.Sprintf("401 %s :No such nick/channel", target)
	}
}
