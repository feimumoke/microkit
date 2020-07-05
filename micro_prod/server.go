package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"zhuhui.com/microkit/micro_pb/pb"
	"zhuhui.com/microkit/micro_prod/pservice"
)

/**
go run server.go --server_address :8001
 */
func main() {
	consulreg := consul.NewRegistry(registry.Addrs("192.168.100.26:8500"))
	service := micro.NewService(
		micro.Name("prodservice"),
		micro.Address(":8100"),
		micro.Registry(consulreg),
	)
	//service.Init()
	micro_grpc_pb.RegisterProdServiceHandler(service.Server(),new(pservice.ProdServiceImpl))
	service.Run()
}
