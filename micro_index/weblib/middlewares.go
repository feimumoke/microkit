package weblib

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"zhuhui.com/microkit/micro_pb/pb"
)

func InitMiddleware(ps micro_grpc_pb.ProdService) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Keys = make(map[string]interface{})
		context.Keys["prodservice"] = ps
		context.Next()
	}
}

func ErrorMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				context.JSON(500, gin.H{"status": fmt.Sprintf("%s", r)})
				context.Abort()
			}
		}()
		context.Next()
	}
}
