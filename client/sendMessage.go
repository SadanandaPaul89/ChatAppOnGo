package main

import (
	"bufio"
	"fmt"
	"os"
)

func (c *Client) sendMessage() {
	string1 := bufio.NewReader(os.Stdin)
	// reader := bufio.NewReader(conn)
	fmt.Print("Enter text: ")
	for {
		fmt.Print("Enter text: ")
		msg, _ := string1.ReadString('\n')
		fmt.Println(c.Username, ":", msg)
		c.Conn.Write([]byte(msg))
	}
}
