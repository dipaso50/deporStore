package cmd

import (
	"fmt"
	"net"
	"testing"
)

type serviceMock struct{}

var reportCalled bool
var productRegistre int

func (sm serviceMock) RegisterProduct(product string) {
	productRegistre++
}

func (sm serviceMock) LimitReached() bool {
	return false
}
func (sm serviceMock) AcceptConnection() {
}
func (sm serviceMock) Report() {
	reportCalled = true
}

func TestConnTerminate(t *testing.T) {
	message := "terminate\n"

	port := 3000

	go func() {

		createClientAndSendMessage(message, t, port)
	}()

	sMock := serviceMock{}

	sen := NewSocketEntry(sMock)

	sen.ServeAndListen(port, 5)

	sen.WaitForIt()

	if !reportCalled {
		t.Errorf("Report method dont called on finish !!")
	}
}

func createClientAndSendMessage(msg string, t *testing.T, port int) {
	conn, err := net.Dial("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		t.Error(err)
	}
	defer conn.Close()

	if _, err := fmt.Fprintf(conn, msg); err != nil {
		t.Error(err)
	}
}

func TestProductRegistration(t *testing.T) {

	productNumber := 3
	port := 4000

	go func() {

		for i := 0; i < productNumber; i++ {
			createClientAndSendMessage("prd", t, port)
		}

		createClientAndSendMessage("terminate\n", t, port)

	}()

	sMock := serviceMock{}

	sen := NewSocketEntry(sMock)

	sen.ServeAndListen(port, 5)

	sen.WaitForIt()

	if productRegistre != productNumber {
		t.Errorf("Expected %d product registered, got %d", productNumber, productRegistre)
	}
}
