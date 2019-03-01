GO ?= go
BIN ?= release
APP ?= gojson


clean:
	rm -rf $(BIN)
	go clean -r . 
	go clean -testcache
	mkdir -p $(BIN)

install:
	go install 

linux: clean
	GOOS=linux GOARCH=amd64 go build -o $(BIN)/$(APP)-linux-amd64 .

darwin: clean
	GOOS=darwin GOARCH=amd64 go build -o $(BIN)/$(APP)-darwin-amd64 .


release: linux darwin
.PHONY: release linux darwin clean install
