
.PHONY: all test clean

all: clean test

build-all: build-all-darwin build-all-windows build-all-linux

build-all-darwin: build-darwin-arm64-free build-darwin-arm64-pro build-darwin-amd64-free build-darwin-amd64-pro

build-all-windows: build-windows-amd64-free build-windows-amd64-pro

build-all-linux: build-linux-amd64-free build-linux-amd64-pro

build-darwin-arm64-free:
	GOOS=darwin GOARCH=arm64 go build -tags darwin -o bin/free/darwin/arm/metaudio ./cmd/cli
	chmod +x bin/metaudio

build-darwin-arm64-pro:
	GOOS=darwin GOARCH=arm64 go build -tags "darwin pro" -o bin/pro/darwin/arm/metaudio ./cmd/cli
	chmod +x bin/metaudio

build-darwin-amd64-free:
	GOOS=darwin GOARCH=amd64 go build -tags darwin -o bin/free/darwin/amd/metaudio ./cmd/cli
	chmod +x bin/metaudio

build-darwin-amd64-pro:
	GOOS=darwin GOARCH=amd64 go build -tags "darwin pro" -o bin/pro/darwin/amd/metaudio ./cmd/cli
	chmod +x bin/metaudio

build-windows-amd64-free:
	GOOS=windows GOARCH=amd64 go build -tags windows -o bin/free/windows/amd/metaudio.exe ./cmd/cli

build-windows-amd64-pro:
	GOOS=windows GOARCH=amd64 go build -tags "windows pro" -o bin/pro/windows/amd/metaudio.exe ./cmd/cli	

build-linux-amd64-free:
	GOOS=linux GOARCH=amd64 go build -tags linux -o bin/free/linux/amd/metaudio ./cmd/cli
	chmod +x bin/metaudio

build-linux-amd64-pro:
	GOOS=linux GOARCH=amd64 go build -tags "linux pro" -o bin/pro/linux/amd/metaudio ./cmd/cli
	chmod +x bin/metaudio

install-darwin-free: 
	go install -tags "darwin" github.com/apolo96/metaudio/cmd/cli/cmd/cli
	mv $(GOPATH)/bin/cli $(GOPATH)/bin/metaudio

install-darwin-pro: 
	go install -tags "darwin pro" github.com/apolo96/metaudio/cmd/cli
	mv $(GOPATH)/bin/cli $(GOPATH)/bin/metaudiopro

install-linux-free: 
	go install -tags "linux" github.com/apolo96/metaudio/cmd/cli
	mv $(GOPATH)/bin/cli $(GOPATH)/bin/metaudio

install-linux-pro: 
	go install -tags "linux pro" github.com/apolo96/metaudio/cmd/cli
	mv $(GOPATH)/bin/cli $(GOPATH)/bin/metaudiopro

install-windows-free: 
	go install -tags "windows" github.com/apolo96/metaudio/cmd/cli
	mv $(GOPATH)/bin/cli.exe $(GOPATH)/bin/metaudio.exe

install-windows-pro: 
	go install -tags "windows pro" github.com/apolo96/metaudio/cmd/cli
	mv $(GOPATH)/bin/cli.exe $(GOPATH)/bin/metaudiopro.exe

test-verbose:
	go test -v ./cmd -tags pro

clean:
	go clean -cache -testcache -modcache
	rm -rf bin/