package main

import (
	"bufio"
	"log"
	"net"
)

func handleConnection(conn net.Conn) {
	log.Println("Incomming connection: ", conn)
	defer conn.Close()

	client := &Client{conn: conn, ch: make(chan string)}

	go writeToClient(client)

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		handleCommand(client, line)
	}

	log.Println("Closing connection: ", conn)
	mutex.Lock()

	// delete the client
	delete(clients, client.username)

	// remove the client from all the channels
	for channel := range channels {
		delete(channels[channel], client)
	}

	mutex.Unlock()
}
