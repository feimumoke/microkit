#!/bin/bash

export MICRO_REGISTRY="etcd"
export MICRO_REGISTRY_ADDRESS="192.168.100.26:2379"
export MICRO_API_NAMESPACE="test.tiger.com"
export MICRO_API_HANDLER="rpc"

echo "MICRO_REGISTRY: " $MICRO_REGISTRY

micro api