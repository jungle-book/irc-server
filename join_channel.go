package main

import "log"

func handleJoinCommand(client *Client, msg IRCMessage) {
	if len(msg.Params) == 0 {
		return
	}

	channelName := msg.Params[0]
	_, ok := channels[channelName]
	if !ok {
		// if the channel dosen't exists create a new channel
		channels[channelName] = make(map[*Client]bool)
	}

	channels[channelName][client] = true
	log.Println("Client joined the channel", channelName, client)
}
