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

	fmt.Printf("Connection established with serveron port %d \n", port)

	var sku string

	for sku != "q" && sku != "terminate" {

		fmt.Printf("Type a sku and hit enter to send it to the server. 'q' or 'terminate' for quit!! \n")

		fmt.Scanln(&sku)

		fmt.Printf("sku :(%s)\n", sku)

		msg := fmt.Sprintf("%s\n", sku)

		if _, err := fmt.Fprintf(conn, msg); err != nil {
			fmt.Errorf("Error sending the message %s %v\n", "test", err)
		}
	}
}
