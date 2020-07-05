package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"zhuhui.com/microkit/grpc_client/helper"
	"zhuhui.com/microkit/grpc_pb/pb"
)

func main() {

	conn, err := grpc.Dial(":8081", grpc.WithTransportCredentials(helper.GetClientCreds()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	prodCli := proc_pb_pb.NewProdServiceClient(conn)
	prodResponse, err := prodCli.GetProdStock(context.Background(), &proc_pb_pb.ProdRequest{ProdId: 12})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(prodResponse.ProdStock)

	responseList, err := prodCli.GetProdStockList(context.Background(), &proc_pb_pb.QuerySize{Size: 4})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(responseList.Prodres)
	fmt.Println("--end--")
}
