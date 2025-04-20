package main

import "net"

type Client struct {
	conn net.Conn
	nick string
	ch chan string
}