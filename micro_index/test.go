package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	myhttp "github.com/micro/go-plugins/client/http"
	"github.com/micro/go-plugins/registry/consul"
	"io/ioutil"
	"log"
	"net/http"
	"zhuhui.com/microkit/micro_pb/pb"
)

func callPluginAPI(s selector.Selector) {
	newClient := myhttp.NewClient(client.Selector(s), client.ContentType("application/json"))
	request := newClient.NewRequest("prodservice", "/v1/prods", micro_pb.ProdsRequest{Size:7})
	var resp micro_pb.ProdListResponse
	err := newClient.Call(context.Background(), request, &resp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.GetData())
}
func callAPI(addr string, path string, method string) (string, error) {
	request, _ := http.NewRequest(method, "http://"+addr+path, nil)
	client := http.DefaultClient
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	buf, _ := ioutil.ReadAll(response.Body)
	return string(buf), nil
}

func main() {
	consulreg := consul.NewRegistry(registry.Addrs("192.168.100.26:8500"))
	//services, err := consulreg.GetService("prodservice")
	//if err != nil {
	//	log.Fatal(err)
	//}

	mySelector := selector.NewSelector(selector.Registry(consulreg), selector.SetStrategy(selector.RoundRobin))
	callPluginAPI(mySelector)
	//for {
	//	next := selector.Random(services)
	//	node, err := next()
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	res, err := callAPI(node.Address, "/v1/prods", "GET")
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	fmt.Println(res)
	//	fmt.Println(node.Id, " || ", node.Address, " || ", node.Metadata)
	//	time.Sleep(time.Second)
	//}

}
