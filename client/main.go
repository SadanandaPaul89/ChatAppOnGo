package main

import (
	"fmt"
	"net"
)

type Client struct {
	Conn     net.Conn
	Username string
	Room     string
}

func main() {
	conn, err := net.Dial("tcp", "localhost:3001")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	client := &Client{
		Conn: conn,
		
	}
	defer conn.Close()
	fmt.Println("Connected to the server.")
	client.username()
	
	go client.receiveMessage()
	
	client.sendMessage()

}
