GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOFMT=$(GOCMD)fmt
BINARY_NAME=$(GOPATH)/bin/citta-server
BINARY_UNIX=$(BINARY_NAME)_unix

build:
	@echo "=============Building Citta Server============="
	$(GOBUILD) -o $(BINARY_NAME) -v

run:
	@echo "=============Running Citta Server============="
	go run main.go

test:
	@echo "=============Running Citta Server Tests============="
	$(GOTEST) -v ./...

clean: 
	@echo "=============Removing Citta Server============="
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

fmt:
	@echo "=============Running gofmt for all project files============="
	$(GOFMT) -w */*.go

docker-build:
	@echo "=============Building Local Citta Server Docker Image============="
	docker build -f ./Dockerfile -t citta-server .

docker-run: docker-build
	@echo "=============Starting Citta-Server Container============="
	docker-compose up -d

docker-stop:
	@echo "=============Stopping Citta-Server Container============="
	docker-compose down

docker-logs:
	@echo "=============Getting Citta-Server Logs============="
	docker-compose logs -f

docker-shell:
	@echo "=============Accessing Container Shell============="
	docker exec -t citta-server bash

docker-clean: docker-down
	@echo "=============Cleaning up============="
	rm -f citta-server
	docker system prune -f
	docker volume prune -f