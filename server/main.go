package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Starting Chat Server on port 3001...")
	listener, err := net.Listen("tcp", ":3001")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Server started successfully. Listening for connections...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}

}
