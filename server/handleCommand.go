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
	}

}
