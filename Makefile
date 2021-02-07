# init go modules
init:
	go mod tidy

# add to vendor
vendor: init
	go mod vendor

# build golang proto compiler
protoc-gen-go:
	go get -u github.com/golang/protobuf/protoc-gen-go

# build proto file.
compile: vendor protoc-gen-go
	protoc -I intern/ --go_out=plugins=grpc:intern/  intern/arash.proto

build: vendor
	go build -o bin/arash  main.go

# clean all build outputs
clean:
	rm -rf bin/
	rm -rf build/
	rm -rf pkg/
	go clean -cache
 
.PHONY: init clean compile build
.DEFAULT_GOAL:=build
