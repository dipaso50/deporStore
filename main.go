package main

import (
	"deportStore/application/feeder"
	"deportStore/infraestructure/cmd"
	"deportStore/infraestructure/repo/inmemory"
	"time"
)

const (
	port              = 4000
	maxClientsAllowed = 5
	timerlimit        = 60 * time.Second
)

func main() {
	immRepo := inmemory.NewInmemoryRepo()
	feederService := feeder.NewFeederService(immRepo)

	socketEntry := cmd.NewSocketEntry(feederService, timerlimit)
	socketEntry.ServeAndListen(port, maxClientsAllowed)
}
