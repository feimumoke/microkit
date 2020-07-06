#!/bin/bash

protoc --micro_out=../pb --go_out=../pb test.proto
protoc-go-inject-tag -input=../pb/test.pb.go