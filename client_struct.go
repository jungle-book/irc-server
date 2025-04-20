package main

import "net"

type Client struct {
	conn net.Conn
	nick string
	username string
	ch chan string
}