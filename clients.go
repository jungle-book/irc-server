package main

import (
	"net"
	"sync"
)

var (
	clients = make(map[net.Conn]*Client)
	mutex sync.Mutex
)