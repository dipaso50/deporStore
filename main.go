package main

import (
	"deportStore/application/feeder"
	"deportStore/infraestructure/cmd"
)

const port = 4000

func main() {
	service := feeder.NewFeederService()
	socketEntry := cmd.NewSocketEntry(service)

	socketEntry.ServeAndListen(port)
}
