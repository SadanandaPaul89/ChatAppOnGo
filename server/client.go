package main

import "net"

type Client struct {
	Conn     net.Conn
	Username string
	Addr     string
	Room     string
}

var clients = make(map[string]*Client)
