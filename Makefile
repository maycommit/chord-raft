PROJECT_NAME := "sdproject"
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
CMD_PATH=cli/main.go
DIST_MAC=dist/mac
DIST_LINUX=dist/linux
BINARY_NAME=sdproject_cli

build:
		mkdir -p $(DIST_MAC) $(DIST_LINUX)
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o ./$(DIST_LINUX)/$(BINARY_NAME) -v $(CMD_PATH)
		GOOS=darwin GOARCH=amd64 $(GOBUILD) -o ./$(DIST_MAC)/$(BINARY_NAME) -v $(CMD_PATH)

proto:
		protoc -I protos/ protos/*.proto --go_out=plugins=grpc:protos