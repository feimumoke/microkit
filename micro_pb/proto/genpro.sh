#!/bin/bash

protoc --micro_out=../pb --go_out=../pb Models.proto
protoc --micro_out=../pb --go_out=../pb ProdService.proto
protoc-go-inject-tag -input=../pb/Models.pb.go
protoc-go-inject-tag -input=../pb/ProdService.pb.go