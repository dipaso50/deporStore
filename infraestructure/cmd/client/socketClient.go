package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Printf("Missing por parameter, usage %s 4000\n", os.Args[0])
		os.Exit(1)
	}

	port, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Printf("Missing por parameter, usage %s 4000\n", os.Args[0])
		os.Exit(1)
	}

	conn, err := net.Dial("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		fmt.Errorf("Error connecting to server %s\n", err)
		os.Exit(2)
	}

	defer func() {
		fmt.Println("Closing client connection")
		conn.Close()
	}()

	if _, err := fmt.Fprintf(conn, "test"); err != nil {
		fmt.Errorf("Error sending the message %s %v\n", "test", err)
	}

}
