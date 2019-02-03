GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOFMT=$(GOCMD)fmt
BINARY_NAME=$(GOPATH)/bin/citta-server
BINARY_UNIX=$(BINARY_NAME)_unix

build:
	@echo "Building Citta Server"
	$(GOBUILD) -o $(BINARY_NAME) -v

run:
	@echo "Building Citta Server"
	go run main.go

test:
	@echo "Running Citta Server Tests"
	$(GOTEST) -v ./...

clean: 
	@echo "Removing Citta Server"
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

fmt:
	@echo "Running gofmt for all project files"
	$(GOFMT) -w */*.go

coverage:
	@echo "Running coverage via Coveralls. It expects you to have set COVERALLS_S3TOOLS_KEY env with coveralls key."
	$(GOCMD) get github.com/mattn/goveralls
	goveralls -repotoken $(COVERALLS_S3TOOLS_KEY)
