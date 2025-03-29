package handlers

import (
	"fmt"
	"log"
	"net"
	"sync"
	"GoChirp/models"
)

// Global map to store messages
var UserMessagesMap = make(map[string]*models.UserMessages)
var mu sync.Mutex // Mutex to handle concurrent access

func HandleConnections(conn net.Conn) {
	defer conn.Close()

	// Read data from connection
	buffer := make([]byte, 5*1024) // 5kb buffer
	var username string

	// First message should be the username
	n, err := conn.Read(buffer)
	if err != nil {
		log.Printf("Read error: %v", err)
		return
	}
	username = string(buffer[:n])

	// Lock mutex, initialize UserMessages for the username, then unlock.
	mu.Lock()
	UserMessagesMap[username] = &models.UserMessages{Username: username, Message: []string{}}
	mu.Unlock()

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			log.Printf("Read error: %v", err)
			return
		}

		// Print the recieved message
		msg := string(buffer[:n])
		fmt.Printf("Recieved: %s", msg)

		mu.Lock()
		UserMessagesMap[username].Message = append(UserMessagesMap[username].Message, msg)
		mu.Unlock()

		// Echo back to the client
		_, err = conn.Write([]byte("Echo:" + msg))
		if err != nil {
			log.Printf("Error Sending Message: %v", err)
			return
		}
	}

}