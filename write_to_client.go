package main

import (
	"fmt"
	"log"
)

func writeToClient(client *Client) {
	for msg := range client.ch {
		log.Println(msg)

		fmt.Fprintf(client.conn, "%s\r\n", msg)
	}
}
