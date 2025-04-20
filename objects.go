package main

import (
	"net"
	"sync"
)

var (
	clients = make(map[net.Conn]*Client)
	channels = make(map[string](map[*Client]bool))
	mutex sync.Mutex
)