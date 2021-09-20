exeName=feeder
sourceName=main.go

all: compileLinux compileWindows

compileLinux:
	go build -o bin/linux/$(exeName) $(sourceName)

compileWindows:
	GOOS=windows GOARCH=386 go build -o bin/windows/$(exeName).exe $(sourceName)

clean:
	rm -f bin/linux/$(exeName) ; rm -f bin/windows/$(exeName).exe ; rm -rf bin/release

test:
	go test -v -cover ./... -count=1 -timeout 90s
	