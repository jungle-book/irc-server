package main

import (
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":6667")
	if err != nil {
		log.Fatal(err)
	}

	defer ln.Close()

	log.Println("Starting server on :6667")

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}

		go handleConnection(conn)
	}
}