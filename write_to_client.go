package main

import (
	"fmt"
)

func writeToClient(client *Client) {
	for msg := range client.ch {
		fmt.Fprintf(client.conn, "%s\r\n", msg)
	}
}
