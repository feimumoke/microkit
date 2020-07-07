package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-micro/service/grpc"
	_ "zhuhui.com/microkit/micro/appinit"
	"zhuhui.com/microkit/micro/pb"
	"zhuhui.com/microkit/micro/service"
)

/**
jsonrpc
apigw.sh
*/
func microService() {

	etcdReg := etcd.NewRegistry(registry.Addrs("192.168.100.26:2379"))

	myService := micro.NewService(
		micro.Name("test.tiger.com.test"),
		micro.Address(":8001"),
		micro.Registry(etcdReg),
	)

	micro_service.RegisterTestServiceHandler(myService.Server(), new(service.TestServiceImpl))
	micro_service.RegisterUserServiceHandler(myService.Server(), new(service.UserServiceImpl))
	myService.Run()

}

func grpcService() {
	etcdReg := etcd.NewRegistry(registry.Addrs("192.168.100.26:2379"))

	myService := grpc.NewService(
		micro.Name("test.tiger.com.test"),
		micro.Address(":8001"),
		micro.Registry(etcdReg),
	)

	micro_service.RegisterTestServiceHandler(myService.Server(), new(service.TestServiceImpl))
	myService.Run()
}

func main() {

	microService()
}
