package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	microHttp "github.com/micro/go-plugins/client/http"
	"log"
)

func main() {
	etcdReg := etcd.NewRegistry(registry.Addrs("192.168.100.26:2379"))
	mySelector := selector.NewSelector(selector.Registry(etcdReg),
		selector.SetStrategy(selector.RoundRobin))
	myClient := microHttp.NewClient(client.Selector(mySelector), client.ContentType("application/json"))

	req := myClient.NewRequest("api.tiger.com.test", "/v1/users", map[string]string{})

	var rsp map[string]interface{}

	err := myClient.Call(context.Background(), req, &rsp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rsp)

}
