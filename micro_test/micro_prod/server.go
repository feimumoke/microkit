package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
	"net/http"
	"zhuhui.com/microkit/micro_test/micro_prod/helper"
	"zhuhui.com/microkit/micro_test/micro_prod/pservice"
)

/**
go run server.go --server_address :8001
 */
func main() {
	consulreg := consul.NewRegistry(registry.Addrs("192.168.100.26:8500"))
	ginrouter := gin.Default()

	v1Group := ginrouter.Group("/v1")
	{
		v1Group.Handle("POST", "/prods", func(context *gin.Context) {
			var pr helper.ProdsRequest
			err := context.Bind(&pr)
			if err != nil || pr.Size <= 0 {
				pr = helper.ProdsRequest{Size: 4}
			}
			context.JSON(http.StatusOK, gin.H{"data": pservice.NewProdList(pr.Size)})
		})
	}

	service := web.NewService(
		//web.Address(":8080"),
		web.Name("prodservice"),
		web.Handler(ginrouter),
		web.Registry(consulreg),
	)
	//service.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	//	writer.Write([]byte("hello micro"))
	//})
	service.Init()
	service.Run()
}
