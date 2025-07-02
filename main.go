package main

import (
	"log/slog"
	"net"
	"os"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger.Info("Starting up the Master Process!")

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		logger.Error("Failed to listen to socket")
		panic(err)
	}
	defer listener.Close()

	logger.Info("Listening on port 8080")
	for {
		conn, err := listener.Accept()
		if err != nil {
			logger.Error("Failed to accept connection")
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger.Info("Accepting connection from: ", "address", conn.RemoteAddr().String())

	buffer := make([]byte, 1024)
	bytes_read, err := conn.Read(buffer)
	if err != nil {
		logger.Error("Failed to read request")
		return
	}
	data := buffer[:bytes_read]

	_, err = conn.Write(data)
	if err != nil {
		logger.Error("Failed to write response")
		return
	}

	logger.Info("Received: ", "data", string(data))
}
