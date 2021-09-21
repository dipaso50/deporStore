package cmd

import (
	"fmt"
	"net"
	"testing"
	"time"
)

var (
	reportCalled        bool
	productRegistre     int
	defaultTimeout      = 60 * time.Second
	defaultClientNumber = 5
)

type serviceMock struct{}

func (sm serviceMock) RegisterProduct(product string) {
	productRegistre++
}

func (sm serviceMock) PrintReport() {
	reportCalled = true
}

func TestConnTerminate(t *testing.T) {
	message := "terminate\n"

	port := 3000

	go func() {
		time.Sleep(1 * time.Second) //esperamos un poco al server
		createClientAndSendMessage(message, t, port)
	}()

	sMock := serviceMock{}

	sen := NewSocketEntry(sMock, defaultTimeout, port, defaultClientNumber)

	sen.ServeAndListen()

	if !reportCalled {
		t.Errorf("Report method dont called on finish !!")
	}
}

func TestProductRegistration(t *testing.T) {

	prds := []string{"prd1\n", "prd2\n", "prd3\n", "terminate\n"}
	productNumber := len(prds) - 1 //restamos uno porque el terminate no debe contar como producto
	port := 6000

	go func() {
		time.Sleep(1 * time.Second) //esperamos un poco al server
		createClientAndSendAllMessages(prds, t, port)
	}()

	sMock := serviceMock{}

	sen := NewSocketEntry(sMock, 6*time.Second, port, defaultClientNumber)

	sen.ServeAndListen()

	if productRegistre != productNumber {
		t.Errorf("Expected %d products, got %d", productNumber, productRegistre)
	}
}

func TestClientLimit(t *testing.T) {
	message := "terminate\n"

	port := 2000
	clientLimit := 1

	go func() {
		time.Sleep(1 * time.Second) //esperamos un poco al server
		createClientAndSendMessage("randommsg", t, port)
		createClientAndSendMessage(message, t, port)
	}()

	sMock := serviceMock{}

	sen := NewSocketEntry(sMock, 5*time.Second, port, clientLimit)

	sen.ServeAndListen()

	if !reportCalled {
		t.Errorf("Report method dont called on finish !!")
	}
}

func createClientAndSendMessage(msg string, t *testing.T, port int) {
	service := fmt.Sprintf(":%d", port)
	conn, err := net.Dial("tcp", service)

	fmt.Printf("Connection to %s\n", service)

	if err != nil {
		t.Errorf("Error creating connection %v", err)
	}

	defer func() {
		if conn != nil {
			conn.Close()
		}
	}()

	time.Sleep(2 * time.Second)

	if _, err := fmt.Fprintf(conn, msg); err != nil {
		t.Error(err)
	}
}

func createClientAndSendAllMessages(allMsg []string, t *testing.T, port int) {
	service := fmt.Sprintf(":%d", port)
	conn, err := net.Dial("tcp", service)

	fmt.Printf("Connection to %s\n", service)

	if err != nil {
		t.Error(err)
	}
	defer conn.Close()

	for _, msg := range allMsg {
		time.Sleep(1 * time.Second)
		fmt.Printf("Sending %s\n", msg)
		if _, err := fmt.Fprintf(conn, msg); err != nil {
			t.Error(err)
		}
	}

}
