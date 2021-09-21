package cmd

import (
	"bufio"
	"deportStore/application/feeder"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

const (
	terminationMessage = "terminate"
)

var (
	chTimeout             = make(chan bool)
	chExplicitTermination = make(chan bool)
	chControlC            = make(chan os.Signal)
)

type SocketEntry struct {
	feederService  feeder.IFeederService
	allConnections []net.Conn
	timeout        time.Duration
}

func NewSocketEntry(s feeder.IFeederService, timeout time.Duration) SocketEntry {
	return SocketEntry{feederService: s, allConnections: make([]net.Conn, 0), timeout: timeout}
}

func (se *SocketEntry) ServeAndListen(port, maxClients int) {

	service := fmt.Sprintf(":%d", port)

	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	se.SetShutdownTimer()

	fmt.Printf("Application started, listening at port %d\n", port)

	go func() {
		for i := 0; i < maxClients; i++ {
			conn, err := listener.Accept()
			if err == nil {
				go se.handleRequest(conn)
			}
		}
	}()

	signal.Notify(chControlC, os.Interrupt, syscall.SIGTERM)

	select {
	case <-chTimeout:
		se.gracefullShutdown()
		return
	case <-chExplicitTermination:
		se.gracefullShutdown()
		return
	case <-chControlC:
		se.gracefullShutdown()
		return
	}

}

func (se *SocketEntry) SetShutdownTimer() {
	timer1 := time.NewTimer(se.timeout)

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
	se.feederService.PrintReport()
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
