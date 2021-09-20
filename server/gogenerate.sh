#!/bin/bash

sudo apt install -y protobuf-compiler libprotobuf-dev
go get google.golang.org/protobuf/cmd/protoc-gen-go google.golang.org/grpc/cmd/protoc-gen-go-grpc
export PATH=$PATH:~/go/bin
go generate . ./proto

