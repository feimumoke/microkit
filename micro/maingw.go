package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"zhuhui.com/microkit/micro/pbgw"
)

/**

降低protoc-gen-go的具体办法，在终端运行如下命令，这里降低到版本 v1.2.0

GIT_TAG="v1.2.0"
go get -d -u github.com/golang/protobuf/protoc-gen-go
git -C "$(go env GOPATH)"/src/github.com/golang/protobuf checkout $GIT_TAG
go install github.com/golang/protobuf/protoc-gen-go

*/
func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	grpcEndpoint := "localhost:8001"
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := micro_service_gw.RegisterTestServiceHandlerFromEndpoint(ctx, mux, grpcEndpoint, opts)
	if err != nil {
		log.Fatal(err)
	}
	http.ListenAndServe(":9000", mux)
}
