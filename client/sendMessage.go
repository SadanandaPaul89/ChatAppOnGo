package main

import (
	"bufio"
	"fmt"
	"os"
)

func (c *Client) sendMessage() {
	string1 := bufio.NewReader(os.Stdin)
	// reader := bufio.NewReader(conn)
	for {
		fmt.Print("Enter text: ")
		msg, _ := string1.ReadString('\n')
		fmt.Println("Sending message:", msg)
		c.Conn.Write([]byte(msg))
	}
}
