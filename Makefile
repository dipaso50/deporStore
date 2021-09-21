exeName=feeder
clientName=clientFeeder
sourceName=main.go

all: buildClient compileLinux  

compileLinux:
	go build -o bin/linux/$(exeName) $(sourceName)

clean:
	rm -f bin/linux/$(exeName) ; rm -f bin/windows/$(exeName).exe ; rm -rf bin/release ; rm -f bin/linux/client/*

test:
	go test  -cover ./... -count=1

testv:
	go test  -v -cover ./... -count=1

runServer: compileLinux
	bin/linux/feeder

runClient: buildClient
	bin/linux/client/clientFeeder 4000

buildClient:
	go build -o bin/linux/client/$(clientName) infraestructure/cmd/client/socketClient.go
	