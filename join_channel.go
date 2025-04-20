package main

import "log"

func joinChannel(client *Client, channelName string) {
	_, ok := channels[channelName]
	if !ok {
		// if the channel dosen't exists create a new channel
		channels[channelName] = make(map[*Client]bool)
	}

	channels[channelName][client] = true 
	log.Println("Client joined the channel", channelName, client)
}