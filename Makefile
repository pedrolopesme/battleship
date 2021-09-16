GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOFMT=$(GOCMD)fmt
BINARY_NAME=$(GOPATH)/bin/battleship-server
BINARY_UNIX=$(BINARY_NAME)_unix

build:
	@echo "=============Building battleship Server============="
	$(GOBUILD) -o $(BINARY_NAME) -v

run:
	@echo "=============Running battleship Server============="
	go run main.go

test:
	@echo "=============Running battleship Server Tests============="
	$(GOTEST) -v ./...

clean: 
	@echo "=============Removing battleship Server============="
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

fmt:
	@echo "=============Running gofmt for all project files============="
	$(GOFMT) -w */*.go

docker-build:
	@echo "=============Building Local battleship Server Docker Image============="
	docker build -f ./Dockerfile -t battleship-server .

docker-run: docker-build
	@echo "=============Starting battleship-Server Container============="
	docker-compose up -d

docker-stop:
	@echo "=============Stopping battleship-Server Container============="
	docker-compose down

docker-logs:
	@echo "=============Getting battleship-Server Logs============="
	docker-compose logs -f

docker-shell:
	@echo "=============Accessing Container Shell============="
	docker exec -t battleship-server bash

docker-clean: docker-down
	@echo "=============Cleaning up============="
	rm -f battleship-server
	docker system prune -f
	docker volume prune -f