package main

import (
	"google.golang.org/grpc"
	"net"
	"zhuhui.com/microkit/grpc_pb/pb"
	"zhuhui.com/microkit/grpc_server/helper"
	"zhuhui.com/microkit/grpc_server/service"
)

func main() {

	creds := helper.GetServerCreds()
	//creds, err := credentials.NewServerTLSFromFile(basepath+"server.crt", basepath+"server.key")

	rpcServer := grpc.NewServer(grpc.Creds(creds))
	proc_pb_pb.RegisterProdServiceServer(rpcServer, new(service.ProdService))
	listener, _ := net.Listen("tcp", ":8081")
	rpcServer.Serve(listener)
	/*
		mux := http.NewServeMux()
		mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
			fmt.Println("request", request)
			rpcServer.ServeHTTP(writer, request)
		})

		_ := &http.Server{
			Addr:    ":8082",
			Handler: mux,
		}
		fmt.Println("start http server")
		httpServer.ListenAndServeTLS(basepath+"server.crt", basepath+"server.key")
	*/
}
