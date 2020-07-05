package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	. "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
	"zhuhui.com/microkit/micro_index/weblib"
	"zhuhui.com/microkit/micro_index/wrappers"
	"zhuhui.com/microkit/micro_pb/pb"
)

type logWapper struct {
	Client
}

func (this *logWapper) Call(ctx context.Context, req Request, rsp interface{}, opts ...CallOption) error {
	fmt.Println("call method")
	md, _ := metadata.FromContext(ctx)
	fmt.Printf("[Log Wapper] ctx: %v service: %s method: %s\n", md, req.Service(), req.Endpoint())
	return this.Client.Call(ctx, req, rsp)
}

func NewLogWapper(c Client) Client {
	return &logWapper{c}
}



func main() {
	consulreg := consul.NewRegistry(registry.Addrs("192.168.100.26:8500"))
	myService := micro.NewService(
		micro.Name("prodservice.client"),
		micro.WrapClient(NewLogWapper),
		micro.WrapClient(wrappers.NewProdsWapper),
	)
	prodService := micro_grpc_pb.NewProdService("prodservice", myService.Client())
	httpServer := web.NewService(
		web.Name("httpprodservice"),
		web.Address(":8001"),
		web.Handler(weblib.NewGinRouter(prodService)),
		web.Registry(consulreg),
	)

	httpServer.Init()
	httpServer.Run()
}
