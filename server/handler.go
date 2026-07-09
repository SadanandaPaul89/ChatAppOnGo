package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
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
	firstMessage, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading from client:", err)
		return
	}
	if strings.HasPrefix(firstMessage, "USERNAME:") {
		client.Username = strings.TrimSpace(strings.TrimPrefix(firstMessage, "USERNAME:"))
	}
	if !strings.HasPrefix(firstMessage, "USERNAME:") {
		fmt.Println("Invalid first message from client:", client.Addr)
		return
	}
	fmt.Printf("***%s has joined the chat***\n", client.Username)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Printf("Client %s disconnected:", client.Username)
				// delete(clients, client.Addr)
				fmt.Println("Clients Remaining:", len(clients))
			} else {
				fmt.Println("Error reading from client:", err)
			}
			return
		}

		fmt.Printf("[%s]: %s", client.Addr, msg)
		broadcastMessage(client, msg)
	}

}
