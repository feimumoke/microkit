#!/bin/bash

protoc --go_out=plugins=grpc:../pb Prod.proto

protoc --grpc-gateway_out=logtostderr=true:../pb Prod.proto