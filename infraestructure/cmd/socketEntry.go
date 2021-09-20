package cmd

import (
	"bufio"
	"deportStore/application/feeder"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

const (
	terminationMessage = "terminate"
	timerlimit         = 60 * time.Second
)

var (
	chTimeout             = make(chan bool)
	chExplicitTermination = make(chan bool)
)

type SocketEntry struct {
	feederService  feeder.IFeederService
	allConnections []net.Conn
}

func NewSocketEntry(s feeder.IFeederService) SocketEntry {
	return SocketEntry{feederService: s, allConnections: make([]net.Conn, 0)}
}

func (se *SocketEntry) ServeAndListen(port, clientLimit int) {

	service := fmt.Sprintf(":%d", port)

	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	SetShutdownTimer()

	fmt.Printf("Application started, listening at port %d\n", port)

	go func() {
		for i := 0; i < clientLimit; i++ {
			conn, err := listener.Accept()
			if err == nil {
				go se.handleRequest(conn)
			}
		}
	}()

}

func (se *SocketEntry) WaitForIt() {

	select {
	case <-chTimeout:
		se.gracefullShutdown()
		return
	case <-chExplicitTermination:
		se.gracefullShutdown()
		return
	}
}

func SetShutdownTimer() {
	timer1 := time.NewTimer(timerlimit)

	go func() {
		<-timer1.C
		fmt.Printf("Timeout !!!\n")
		chTimeout <- true
	}()
}

func (se *SocketEntry) handleRequest(conn net.Conn) {

	feeder := se.feederService
	se.allConnections = append(se.allConnections, conn)

	fmt.Println("Connected with a new client")

	feeder.AcceptConnection()

	for {
		prod, _ := bufio.NewReader(conn).ReadString('\n')

		prod = strings.Replace(prod, "\n", "", -1)

		if len(prod) == 0 {
			continue
		}

		fmt.Printf("Message received (%s)\n", prod)

		if prod == terminationMessage {
			chExplicitTermination <- true
			return
		}

		feeder.RegisterProduct(prod)
	}
}

func (se *SocketEntry) gracefullShutdown() {

	if len(se.allConnections) == 0 {
		fmt.Println("No connections")
		return
	}

	fmt.Printf("Gracefull shutdown, closing %d connections\n", len(se.allConnections))

	for _, con := range se.allConnections {
		con.Close()
	}
	se.feederService.Report()
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
