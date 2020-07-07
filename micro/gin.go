package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"zhuhui.com/microkit/micro/sidecar"
)

/**
gin 集成到go micro
*/

func main() {
	ginRouter := gin.Default()
	v1 := ginRouter.Group("/v1")
	{
		v1.Handle("POST", "/users", func(ginCtx *gin.Context) {
			ginCtx.JSON(200, gin.H{"data": "test"})
		})
	}
	server := http.Server{Addr: ":8088", Handler: ginRouter}

	handler := make(chan error)
	go func() {
		handler <- server.ListenAndServe()
	}()

	go func() {
		notify := make(chan os.Signal)
		signal.Notify(notify, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
		handler <- fmt.Errorf("%s", <-notify)
	}()

	service := sidecar.NewService("api.tiger.com.test")
	fmt.Println(uuid.New().String())
	service.AddNode("test-"+uuid.New().String(), 8088, "192.168.100.26:8088")

	go func() {
		//register
		err := sidecar.RegService(service)
		if err != nil {
			handler <- err
		}
	}()
	getHandler := <-handler
	fmt.Println(getHandler.Error())

	err := sidecar.UnRegService(service)
	if err != nil {
		fmt.Println(err)
	}

	err = server.Shutdown(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
