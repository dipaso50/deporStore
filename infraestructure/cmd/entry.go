package cmd

import (
	"bufio"
	"deportStore/application/feeder"
	"fmt"
	"net"
	"os"
)

type SocketEntry struct {
	feederService feeder.IFeederService
}

func NewSocketEntry(s feeder.IFeederService) SocketEntry {
	return SocketEntry{s}
}

func (se SocketEntry) ServeAndListen(port int) {

	service := fmt.Sprintf(":%d", port)

	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		se.handleRequest(listener)
	}
}

func (se SocketEntry) handleRequest(listener *net.TCPListener) {

	if se.feederService.LimitReached() {

		se.feederService.AcceptConnection()

		go func(lis *net.TCPListener, feeder feeder.IFeederService) {
			conn, err := lis.Accept()

			if err != nil {
				return
			}

			prod, _ := bufio.NewReader(conn).ReadString('\n')

			feeder.RegisterProduct(prod)

			conn.Close() // we're finished with this client
		}(listener, se.feederService)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
