# GoChirp

GoChirp is a simple TCP-based chat application written in Go. It allows real-time, two-way communication between clients and a server. The server listens for incoming connections, processes messages, and echoes them back to the sender.

## Features
- Lightweight TCP server for handling chat messages.
- Concurrent connection handling using goroutines.
- Simple echo functionality for testing communication.

## Installation
1. Clone the repository:
   ```sh
   git clone https://github.com/your-username/GoChirp.git
   cd GoChirp
   ```
2. Build the project:
   ```sh
   go build -o gochirp
   ```
3. Run the server:
   ```sh
   ./gochirp
   ```

## Usage
1. The server listens on `172.19.25.26:8080`.
2. Clients can connect using tools like `nc` (Netcat):
   ```sh
   nc 172.19.25.26 8080
   ```
3. Type messages and see them echoed back.

## File Structure
```
GoChirp/
│── handlers/
│   ├── connections.go  # Handles client connections
│── main.go             # Server entry point
│── README.md           # Project documentation
```

## Future Enhancements
- Implement user authentication.
- Support multiple clients with message broadcasting.
- Add a GUI-based client.

## License
MIT License
