GO ?= go
BIN ?= bin
APP ?= gojson



clean:
	rm -rf bin
	go clean -r . 
	go clean -testcache

install:
	go install 

build: clean
	mkdir -p $(BIN)
	go build -o $(BIN)/$(APP) .

