package main

import (
	"fmt"
	"strings"
)

// find client by username and send a private message to them
func findClientByUsername(username string) (*Client, bool) {
	clientsMutex.RLock()
	defer clientsMutex.RUnlock()
	for _, client := range clients {
		if client.Username == username {

			return client, true
		}
	}
	return nil, false
}

func handleCommand(client *Client, msg string) {

	parts := strings.Fields(msg)

	if len(parts) == 0 {
		return
	}
	switch parts[0] {
	//text a specific user privately
	case "/msg":
		//handle private message
		if len(parts) < 3 {
			client.Conn.Write([]byte("Usage: /msg <username> <message>\n"))
			return
		}
		recipient := parts[1]                          //Extract the recipient username
		privateMessage := strings.Join(parts[2:], " ") //Extract the message

		//Search for recipients
		findClient, found := findClientByUsername(recipient)
		if found {
			//Private message found, send it to the recipient and the sender
			_, err := findClient.Conn.Write([]byte(fmt.Sprintf("[Private --> %s]: %s\n", client.Username, privateMessage)))
			if err != nil {
				fmt.Printf("Error sending private message to %s: %v\n", recipient, err)
				return
			}
			_, err = client.Conn.Write([]byte(fmt.Sprintf("[Private] To %s: %s\n", recipient, privateMessage)))
			if err != nil {
				fmt.Printf("Error sending private message to %s: %v\n", client.Username, err)
				return
			}

		}

		//If the recipient is not found, send an error message back to the sender
		client.Conn.Write([]byte(fmt.Sprintf("User %s not found\n", recipient)))

	case "/join":
		//handle join room
		if len(parts) < 2 {
			client.Conn.Write([]byte("Usage: /join <room>\n"))
			return
		}
		room := parts[1]
		client.Room = room
		client.Conn.Write([]byte(fmt.Sprintf("You have joined room: %s\n", room)))
		broadcastRoomSystem(client, fmt.Sprintf("%s has joined the room: %s\n", client.Username, room))
		return
	}

	//leave room command
	if parts[0] == "/leave" {
		if client.Room == "" {
			client.Conn.Write([]byte("You are not in a room\n"))
			return
		}
		broadcastRoomSystem(client, fmt.Sprintf("%s has left the room: %s\n", client.Username, client.Room))
		client.Room = ""
		client.Conn.Write([]byte("You have left the room\n"))
		return
	}

}
