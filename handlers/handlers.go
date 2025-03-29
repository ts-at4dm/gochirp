package handlers

import (
	"fmt"
	"log"
	"net"
)

func HandleConnections(conn net.Conn) {
	defer conn.Close()

	// Read data from connection
	buffer := make([]byte, 5*1024) // 5kb buffer
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			log.Printf("Read error: %v", err)
			return
		}

		// Print the recieved message
		msg := string(buffer[:n])
		fmt.Printf("Recieved: %s", msg)

		// Echo back to the client
		_, err = conn.Write([]byte("Echo:" + msg))
		if err != nil {
			log.Printf("Error Sending Message: %v", err)
			return
		}
	}

}