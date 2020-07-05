package util

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/consul"
	"github.com/go-kit/kit/sd/lb"
	transhttp "github.com/go-kit/kit/transport/http"
	"github.com/hashicorp/consul/api"
	"io"
	"net/url"
	"os"
	. "zhuhui.com/microkit/kit_client/cs"
)

func GetUser() (string, error) {
	config := api.DefaultConfig()
	config.Address = "192.168.100.26:8500"
	api_cli, err := api.NewClient(config)
	if err != nil {
		return "", err
	}
	client := consul.NewClient(api_cli)
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stdout)
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
			endpoints, err := endpointer.Endpoints()
			if err != nil {
				return "", err
			}
			fmt.Println("endpoint len:", len(endpoints))
			mylb := lb.NewRoundRobin(endpointer)
			for {
				//getUserInfo := endpoints[0]
				getUserInfo, err := mylb.Endpoint()
				if err != nil {
					return "", err
				}
				ctx := context.Background()
				resp, err := getUserInfo(ctx, UserRequest{Uid: 102})
				if err != nil {
					return "", err
				}
				userInfo := resp.(UserResponse)
				return userInfo.Result, nil
			}
		}
	}
}
