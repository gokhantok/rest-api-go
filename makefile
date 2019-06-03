    # Go parameters
    GOCMD=go
    GOBUILD=$(GOCMD) build
    GOCLEAN=$(GOCMD) clean
    BINARY_NAME=myapp
    
    
    all:  docker-build docker-run
    build: 
		$(GOBUILD) -o $(BINARY_NAME) ./cmd/myapp/*.go  
    test: ls
		$(GOTEST) -v ./...
    clean: 
		$(GOCLEAN)
		rm -f $(BINARY_NAME)
    run:
		$(GOCMD) run ./cmd/myapp/*.go   

    protoc:
		protoc --go_out=. ./cmd/myapp/*.proto
	
    docker-build:
		docker build -t webserver .    
    docker-run:
		docker run -p:8000:8000 -it  webserver:latest