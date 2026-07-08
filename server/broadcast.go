package main

import "fmt"

func broadcastMessage(sender *Client, message string) {
	formatted := fmt.Sprintf("Message from %s: %s", sender.Username, message)
	for _, client := range clients {
		if client != sender {
			_, err := client.Conn.Write([]byte(formatted))
			if err != nil {
				fmt.Println("Error broadcasting message to", client.Addr, ":", err)
			}
		}
	}
	

}
