package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/web"
	"net/http"
)

func main() {
	ginrouter := gin.Default()
	ginrouter.Handle("GET", "/", func(context *gin.Context) {
		data := make([]interface{}, 0)
		context.JSON(http.StatusOK, gin.H{"data": data,})
	})

	service := web.NewService(web.Address(":8000"),
		web.Handler(ginrouter),
	)
	service.Run()
}
