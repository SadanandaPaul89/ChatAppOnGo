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
	clientsMutex.Lock()
	clients[client.Addr] = client
	clientsMutex.Unlock()
	defer func() {
		clientsMutex.Lock()
		delete(clients, client.Addr)
		clientsMutex.Unlock()
		fmt.Printf("***%s has left the chat***\n", client.Username)                           //server log
		broadcastSystem(client, fmt.Sprintf("***%s has left the chat***\n", client.Username)) //broadcast to all clients
	}() //anonymous function to remove the client from the map when the connection is closed
	//clientAddr := conn.RemoteAddr().String()
	fmt.Println("Client connected:", client.Username, "from", client.Addr)
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
	fmt.Printf("***%s has joined the chat***\n", client.Username)                           //server log
	broadcastSystem(client, fmt.Sprintf("***%s has joined the chat***\n", client.Username)) //broadcast to all clients
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
		if strings.HasPrefix(msg, "/") {
			handleCommand(client, msg)
			continue
		}

		fmt.Printf("[%s]: %s", client.Username, msg)
		broadcastMessage(client, msg)
	}

}
