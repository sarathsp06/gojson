GO ?= GO111MODULE=on go
BIN ?= release
APP ?= gojson
GOOS ?= darwin
GOARCH ?= amd64

clean: mod
	rm -rf $(BIN)
	$(GO) clean -r . 
	$(GO) clean -cache
	$(GO) clean -testcache
	mkdir -p $(BIN)

mod:
	$(GO) mod tidy -v

install:
	$(GO) install 

linux: clean
	GOOS=linux GOARCH=amd64 $(GO) build -o $(BIN)/$(APP)-linux-amd64 .

darwin: clean
	GOOS=darwin GOARCH=amd64 $(GO) build -o $(BIN)/$(APP) .

build: mod
	GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) build -o $(BIN)/$(APP)-$(GOOS)-$(GOARCH) .

build-all:
	$(GO) tool dist list  | xargs -I {} python -c "import sys;os,arch=sys.argv[1].split('/');print 'GOOS=%s GOARCH=%s make build' %(os,arch)" {} | sh



release: linux darwin
.PHONY: release mod install darwin linux build build-all
