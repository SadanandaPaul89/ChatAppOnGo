package main

import (
	"bufio"
	"fmt"
)

func (c *Client) receiveMessage() {
	reader := bufio.NewReader(c.Conn)
	for {
		new, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from server:", err)
			return
		}
		fmt.Print("\r")
		fmt.Print(new)
		fmt.Print("Enter text: ")
	}

}
