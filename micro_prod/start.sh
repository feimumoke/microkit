#!/bin/bash

gnome-terminal -x go run server.go --server_address :8001
gnome-terminal -x go run server.go --server_address :8002
gnome-terminal -x go run server.go --server_address :8003