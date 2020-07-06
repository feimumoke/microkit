#!/bin/bash

#protoc --micro_out=../pb --go_out=../pb test.proto
#protoc-go-inject-tag -input=../pb/test.pb.go

#protoc --go_out=plugins=grpc:../pbgw test.proto
#protoc --grpc-gateway_out=logtostderr=true:../pbgw test.proto


protoc --go_out=../pb Models.proto
protoc --micro_out=../pb --go_out=../pb UserService.proto
protoc-go-inject-tag -input=../pb/Models.pb.go
protoc-go-inject-tag -input=../pb/UserService.pb.go