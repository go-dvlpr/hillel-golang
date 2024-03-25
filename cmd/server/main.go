package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:3232")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer listener.Close()

	fmt.Println("Server is listening on port 3232")

	for {
		// Accept incoming connections
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		log.Println("somebody connected")

		// Handle client connection in a goroutine
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) error {
	tmp := make([]byte, 256)

	defer conn.Close()

	_, err := conn.Read(tmp)
	if err != nil {
		log.Println("failed to read ", err)
		return err
	}

	fmt.Println(string(tmp))

	return nil
}
