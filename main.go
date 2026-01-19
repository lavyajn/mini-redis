package main

import (
	"fmt"
	"net"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func main() {
	//Create a TCP listerner on port 6379
	listener, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println("Error binding to port:", err)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Println("Server is listening on Port 6379...")

	// we loop forever to accept many connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		// a goroutine to handle the new connections immediately
		go handleconnection(conn)
	}
}

func handleconnection(conn net.Conn) {

	defer conn.Close()

	reader := bufio.NewReader(conn)

	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading from connection:", err)
		return
	}

	line = strings.TrimSpace(line)
	
	if line[0] != '*' {
		fmt.Println("Error: Not a valid RESP array")
		return
	}

	count, err := strconv.Atoi(line[1:])
	if err != nil {
		fmt.Println("Error parsing array length:", err)
		return
	}

	for(i := 0; i < count; ++i) {
		


	}

	fmt.Println("Client sent a command with", count, "arguments")

	/* // buffer to hold incoming data
	buf := make([]byte, 1024)

	//read data from the client
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading from connection:", err)
		return
	}

	//print the received data
	receivedData := string(buf[:n])
	fmt.Println("Received data:", receivedData) */
}
