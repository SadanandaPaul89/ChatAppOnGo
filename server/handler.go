package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	client := &Client{
		Conn:     conn,
		Username: "",
		Addr:     conn.RemoteAddr().String(),
		Room:     "",
	}
	clients[client.Addr] = client
	defer delete(clients, client.Addr)
	//clientAddr := conn.RemoteAddr().String()
	fmt.Println("Client connected:", client.Addr)
	fmt.Println("Clients Connected:", len(clients))
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("Client disconnected:", client.Addr)
				// delete(clients, client.Addr)
				fmt.Println("Clients Remaining:", len(clients))
			} else {
				fmt.Println("Error reading from client:", err)
			}
			return
		}

		fmt.Printf("Received message from %s: %s", client.Addr, msg)
		broadcastMessage(client, msg)
	}

}
