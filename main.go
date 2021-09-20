package main

import (
	"deportStore/application/feeder"
	"deportStore/infraestructure/cmd"
)

const port = 4000

func main() {
	maxClientsAllowed := 5
	service := feeder.NewFeederService(maxClientsAllowed)
	socketEntry := cmd.NewSocketEntry(service)

	socketEntry.ServeAndListen(port)
}
