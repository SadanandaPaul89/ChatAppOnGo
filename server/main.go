package main

import (
	"fmt"
	"net"
)

// type Client struct {
// 	conn     net.Conn
// 	Username string
// 	Addr     string
// 	Room     string
// }

// var clients = make(map[string]*Client)

// func handleConnection(conn net.Conn) {
// 	defer conn.Close()
// 	clientAddr := conn.RemoteAddr().String()
// 	fmt.Println("Client connected:", clientAddr)
// 	reader := bufio.NewReader(conn)
// 	for {
// 		msg, err := reader.ReadString('\n')
// 		if err != nil {
// 			if err == io.EOF {
// 				fmt.Println("Client disconnected:", clientAddr)
// 			} else {
// 				fmt.Println("Error reading from client:", err)
// 			}
// 			return
// 		}

// 		fmt.Printf("Received message from %s: %s", clientAddr, msg)
// 		_, erre := conn.Write([]byte(msg))
// 		if erre != nil {
// 			fmt.Println(erre)
// 			return
// 		}
// 	}

// }

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
