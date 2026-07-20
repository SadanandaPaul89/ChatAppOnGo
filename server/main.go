package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func startTCPSERVER() {
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

func main() {
	//start tcp chat server
	go startTCPSERVER()
	//serve static files
	http.Handle("/", http.FileServer(http.Dir("./web")))
	fmt.Println("Starting web server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
