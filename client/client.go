package main

import (
	"fmt"
	"bufio"
	"net"
	"os"
	"strings"
)

// Establishes a connection to the server at the inputted IP
func serverConnect(ipAddress string) (net.Conn, error) {
	conn, err := net.Dial("tcp", ipAddress)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func main() {
	//Reads line of input
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Print("Enter the IP Address: ")
	ipAddress, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	// Trim whitespace and newline from input
	ipAddress = strings.TrimSpace(ipAddress)
	
	
	// Asks for a username
	fmt.Print("Enter desired username: ")
	username, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading name:", err)
		return
	}
	username = strings.TrimSpace(username)
	
	// Connect to server
	conn, err := serverConnect(ipAddress)
	if err != nil {
		fmt.Println("Error Connecting to server:", err)
		return
	}
	defer conn.Close() //closes connection when done
	
	fmt.Println("Connected, Welcome to GoChirp", username)
}