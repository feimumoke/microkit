#!/bin/bash

protoc --micro_out=../pb --go_out=../pb Prods.proto
protoc-go-inject-tag -input=../pb/Prods.pb.go