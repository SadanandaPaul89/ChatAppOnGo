package main

import (
	"net"
	"sync"
)

type Client struct {
	Conn     net.Conn
	Username string
	Addr     string
	Room     string
}

var clients = make(map[string]*Client)
var clientsMutex sync.RWMutex
