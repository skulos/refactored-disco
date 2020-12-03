GOVERSION=go1.13.4

export PROJDIR := $(shell pwd)
export GOBIN := $(PROJDIR)/bin
export PATH := $(PATH):$(GOBIN)

GO_IN_PATH := $(shell command -v go 2> /dev/null)
ifndef GO_IN_PATH
export PATH := $(PATH):/usr/local/$(GOVERSION)/bin:$(GOBIN)
export GOROOT=/usr/local/$(GOVERSION)
endif

grpcgatewaypath=$(shell go list -m -f '{{.Dir}}' github.com/grpc-ecosystem/grpc-gateway)
googleapis_path=$(grpcgatewaypath)/third_party/googleapis

# init go modules
init:
	go mod tidy

# add to vendor
vendor: init
	go mod vendor

# build golang proto compiler
bin/protoc-gen-go:
	go install -mod=vendor -v github.com/golang/protobuf/protoc-gen-go

# build bfruntime proto file.
bfruntime: init bin/protoc-gen-go
	protoc -I src/bfruntime/ -I ${googleapis_path} --go_out=plugins=grpc:./  src/bfruntime/bfruntime.proto

# build our golang source
install:
	go install -mod=vendor -v ./...

# clean all build outputs
clean:
	rm -rf bin/
	rm -rf build/
	rm -rf pkg/
	go clean -cache
 
.PHONY: init install clean bfruntime
.DEFAULT_GOAL:=install
