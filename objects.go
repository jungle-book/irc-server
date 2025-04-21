package main

import (
	"sync"
)

var (
	clients  = make(map[string]*Client)
	channels = make(map[string](map[*Client]bool))
	mutex    sync.Mutex
)
