package weblib

import (
	"github.com/gin-gonic/gin"
	"zhuhui.com/microkit/micro_pb/pb"
)

func NewGinRouter(ps micro_grpc_pb.ProdService) *gin.Engine {
	ginrouter := gin.Default()
	ginrouter.Use(InitMiddleware(ps),ErrorMiddleware())
	viGroup := ginrouter.Group("/v1")
	{
		viGroup.Handle("POST", "/prods", GetProdsList)
		viGroup.Handle("GET", "/prod/:pid", GetProdDetail)
	}
	return ginrouter
}
