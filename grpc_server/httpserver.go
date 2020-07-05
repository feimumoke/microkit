package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"zhuhui.com/microkit/grpc_pb/pb"
	"zhuhui.com/microkit/grpc_server/helper"
)

func main() {
	mux := runtime.NewServeMux()
	opt := []grpc.DialOption{grpc.WithTransportCredentials(helper.GetClientCreds())}
	err := proc_pb_pb.RegisterProdServiceHandlerFromEndpoint(context.Background(),
		mux, "localhost:8081", opt)
	if err != nil {
		log.Fatal(err)
	}
	hserver := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	hserver.ListenAndServe()
}
