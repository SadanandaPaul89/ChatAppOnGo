package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:3001")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()
	fmt.Println("Connected to the server.")

	string1 := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter text: ")
		msg, _ := string1.ReadString('\n')
		fmt.Println("Sending message:", msg)
		conn.Write([]byte(msg))
	}

}
