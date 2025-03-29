package main

import (
	"fmt"
	"net"
	"log"
	"GoChirp/handlers"
)

func main() {
	// Listen on TCP port 8080
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	defer listener.Close()

	port := listener.Addr().(*net.TCPAddr).Port

	log.Printf("Server Launch Confirmed...Port: %d", port)

	for {
		// Accepts a new connection
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error: %v", err)
			log.Printf("Connection acception error: %v", err)
		}

		// Handle the connection in a new goroutine
		go handlers.HandleConnections(conn)
	}
}
