package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func (c *Client) username(){
	usr:= bufio.NewReader(os.Stdin)
	fmt.Print("Enter your username: ")
	username, _ := usr.ReadString('\n')
	c.Username = username
	c.Username = strings.TrimSpace(c.Username)
	c.Conn.Write([]byte("Username: "+ c.Username + "\n"))
}