package main

import (
	"fmt"
)

func broadcastMessage(sender *Client, message string) {
	var formatted string
	if sender.Username != "" {
		formatted = fmt.Sprintf("[%s]: %s\n", sender.Username, message)
	} else {
		formatted = fmt.Sprintf("[%s]:%s %s\n", sender.Room, sender.Username, message)
	}
	clientsMutex.RLock()
	defer clientsMutex.RUnlock()
	for _, client := range clients {
		if client == sender {
			continue
		}
		//New skip clients in a room that are not in the same room as the sender
		if client.Room != sender.Room {
			continue
		}

		_, err := client.Conn.Write([]byte(formatted))
		if err != nil {
			fmt.Println("Error broadcasting message to", client.Addr, ":", err)
		}

	}

}
