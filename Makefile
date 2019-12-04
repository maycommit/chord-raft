PROJECT_NAME := "sdproject"
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
CMD_PATH := ./cmd/main.go
DIST_MAC := dist/mac
DIST_LINUX := dist/linux

build:
		mkdir -p $(DIST_MAC) $(DIST_LINUX)
		export MODULE=$(GO111MODULE=on go list -m)
		GOOS=linux GOARCH=amd64 $(GOBUILD) -tags release -ldflags '-X $(MODULE)/cmd.Version=$(VERSION) -X $(MODULE)/cmd.BuildDate=$(DATE)' -o ./$(DIST_LINUX)/$(BINARY_NAME) -v $(CMD_PATH)
		GOOS=darwin GOARCH=amd64 $(GOBUILD) -tags release -ldflags '-X $(MODULE)/cmd.Version=$(VERSION) -X $(MODULE)/cmd.BuildDate=$(DATE)' -o ./$(DIST_MAC)/$(BINARY_NAME) -v $(CMD_PATH)

proto:
		protoc -I protos/ protos/*.proto --go_out=plugins=grpc:protos