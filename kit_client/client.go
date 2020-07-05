package main

import (
	"context"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/go-kit/kit/endpoint"
	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/consul"
	"github.com/go-kit/kit/sd/lb"
	transhttp "github.com/go-kit/kit/transport/http"
	"github.com/hashicorp/consul/api"
	"io"
	"log"
	"net/url"
	"os"
	"time"
	. "zhuhui.com/microkit/kit_client/cs"
	"zhuhui.com/microkit/kit_client/util"
)

func directConn() {
	target, _ := url.Parse("http://localhost:8080")
	client := transhttp.NewClient("GET", target, GetUserInfoRequest, GetUserInfoResponse)
	user_endpoint := client.Endpoint()
	ctx := context.Background()
	resp, err := user_endpoint(ctx, UserRequest{Uid: 102})
	if err != nil {
		fmt.Println(err)
	}
	userInfo := resp.(UserResponse)
	fmt.Println(userInfo.Result)
}

func consulConn() {
	{
		config := api.DefaultConfig()
		config.Address = "192.168.100.26:8500"
		api_cli, _ := api.NewClient(config)
		client := consul.NewClient(api_cli)
		var logger kitlog.Logger
		{
			logger = kitlog.NewLogfmtLogger(os.Stdout)
		}
		{
			tags := []string{"primary"}
			instance := consul.NewInstancer(client, logger, "userservice", tags, true)
			{
				f := func(ser_url string) (endpoint.Endpoint, io.Closer, error) {
					tart, _ := url.Parse("http://" + ser_url)
					return transhttp.NewClient("GET", tart, GetUserInfoRequest, GetUserInfoResponse).Endpoint(), nil, nil
				}
				endpointer := sd.NewEndpointer(instance, f, logger)
				endpoints, _ := endpointer.Endpoints()
				fmt.Println("endpoint len:", len(endpoints))
				mylb := lb.NewRoundRobin(endpointer)
				for {
					//getUserInfo := endpoints[0]
					getUserInfo, _ := mylb.Endpoint()
					ctx := context.Background()
					resp, err := getUserInfo(ctx, UserRequest{Uid: 102})
					if err != nil {
						fmt.Println(err)
					}
					userInfo := resp.(UserResponse)
					fmt.Println(userInfo.Result)
					time.Sleep(30 * time.Millisecond)
				}

			}
		}
	}
}

func main() {
	configA := hystrix.CommandConfig{
		Timeout:                2000, //超时时间
		MaxConcurrentRequests:  2,    //最大并发数
		RequestVolumeThreshold: 5,    //熔断器请求阈值，有5个请求才进行错误百分比计算
		ErrorPercentThreshold:  99,
		SleepWindow:            10} //尝试去请求
	hystrix.ConfigureCommand("getuser", configA)
	err := hystrix.Do("getuser", func() error {
		res, err := util.GetUser()
		fmt.Println(res)
		return err
	}, func(e error) error {
		fmt.Println("降级用户")
		return e
	})
	if err != nil {
		log.Fatal(err)
	}
}
