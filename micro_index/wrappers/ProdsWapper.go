package wrappers

import (
	"context"
	"github.com/afex/hystrix-go/hystrix"
	. "github.com/micro/go-micro/client"
	"strconv"
	. "zhuhui.com/microkit/micro_pb/pb"
)

type ProdsWrapper struct {
	Client
}

func newProd(id int32, pname string) *ProdModel {
	return &ProdModel{ProdID: id, ProdName: pname}
}

func fallbackData(rsp interface{})  {
	switch t:=rsp.(type) {
	case *ProdListResponse:
		fallbackProds(rsp)
	case *ProdDetailResponse:
		t.Data=newProd(11,"prod fallback detail")
	}
}
func fallbackProds(rsp interface{}) {
	models := make([]*ProdModel, 0)
	var i int32
	for i = 0; i < 3; i++ {
		models = append(models, newProd(1000+i, "prod-"+strconv.Itoa(1000+int(i))))
	}
	result := rsp.(*ProdListResponse)
	result.Data = models
}

func (this *ProdsWrapper) Call(ctx context.Context, req Request, rsp interface{}, opts ...CallOption) error {
	cmdName := req.Service() + "." + req.Endpoint()
	hConfig := hystrix.CommandConfig{
		Timeout: 1000,
		RequestVolumeThreshold:4,
		ErrorPercentThreshold:50,
		SleepWindow:5000,
	}
	hystrix.ConfigureCommand(cmdName, hConfig)
	return hystrix.Do(cmdName, func() error {
		return this.Client.Call(ctx, req, rsp)
	}, func(e error) error {
		fallbackData(rsp)
		return nil
	})
}

func NewProdsWapper(c Client) Client {
	return &ProdsWrapper{c}
}
