package main

import "fmt"

func listChannels(client *Client) {
	for channel := range channels {
		client.ch <- fmt.Sprint(channel)
	}
}