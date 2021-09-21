exeName=feeder
clientName=clientFeeder
sourceName=main.go

all: buildClient compileLinux compileWindows

compileLinux:
	go build -o bin/linux/$(exeName) $(sourceName)

compileWindows:
	GOOS=windows GOARCH=386 go build -o bin/windows/$(exeName).exe $(sourceName)

clean:
	rm -f bin/linux/$(exeName) ; rm -f bin/windows/$(exeName).exe ; rm -rf bin/release

test:
	go test  -cover ./... -count=1

testv:
	go test  -v -cover ./... -count=1

run:
	go run main.go

runClient:
	go run infraestructure/cmd/client/socketClient.go  4000

buildClient:
	go build -o bin/linux/client/$(clientName) infraestructure/cmd/client/socketClient.go
	