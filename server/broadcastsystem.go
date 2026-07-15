package main

import "fmt"

func broadcastSystem(Sender *Client, message string) {
	clientsMutex.RLock()
	defer clientsMutex.RUnlock()
	for _, client := range clients {
		_, err := client.Conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Error broadcasting message to", client.Addr, ":", err)
		}
	}
}
