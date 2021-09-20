package main

import (
	"deportStore/application/feeder"
	"deportStore/infraestructure/cmd"
)

const (
	port              = 4000
	maxClientsAllowed = 5
)

func main() {

	feederService := feeder.NewFeederService(maxClientsAllowed)
	socketEntry := cmd.NewSocketEntry(feederService)

	socketEntry.ServeAndListen(port, maxClientsAllowed)
	socketEntry.WaitForIt()
}
