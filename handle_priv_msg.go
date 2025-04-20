package main

import (
	"fmt"
	"strings"
)

func handlePrivMessage(client *Client, msg IRCMessage) {
	target := msg.Params[0]

	mutex.Lock()
	defer mutex.Unlock()

	if strings.Contains(msg.Params[0], "#") {
		if members, exists := channels[target]; exists {
			for member := range members {
				if member.username != client.username {
					member.ch <- fmt.Sprintf(":%s PRIVMSG %s :%s", client.nick, target, msg.Trailing)
				}
			}
		}
	} else {
		for _, recipent := range clients {
			if recipent.nick == target {
				recipent.ch <- fmt.Sprintf(":%s PRIVMSG %s :%s", client.nick, target, msg.Trailing)
				return
			}
		}

		client.ch <- fmt.Sprintf("401 %s :No such nick/channel", target)
	}
}
