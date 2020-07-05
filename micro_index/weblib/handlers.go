package weblib

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	. "zhuhui.com/microkit/micro_pb/pb"
)

func PaincIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func GetProdDetail(ginCtx *gin.Context) {
	var prodReq ProdsRequest
	PaincIfError(ginCtx.BindUri(&prodReq))
	prodService := ginCtx.Keys["prodservice"].(ProdService)
	fmt.Println("xxx",prodReq.ProdId)
	response, _ := prodService.GetProdDetail(context.Background(), &prodReq)
	fmt.Println("xxxx",response)
	ginCtx.JSON(http.StatusOK, gin.H{"data": response.Data})
}

func GetProdsList(ginCtx *gin.Context) {
	prodService := ginCtx.Keys["prodservice"].(ProdService)
	var prodReq ProdsRequest
	err := ginCtx.Bind(&prodReq)
	if err != nil {
		ginCtx.JSON(500, gin.H{"status": err.Error()})
	} else {

		response, err := prodService.GetProdsList(context.Background(), &prodReq)
		if err != nil {
			ginCtx.JSON(500, gin.H{"status": err.Error()})
		} else {
			ginCtx.JSON(http.StatusOK, gin.H{"data": response.Data})
		}

		/*
		hConfig := hystrix.CommandConfig{
			Timeout: 2000,
		}
		hystrix.ConfigureCommand("getprods", hConfig)
		err := hystrix.Do("getprods", func() error {
			response, err = prodService.GetProdsList(context.Background(), &prodReq)
			if err != nil {
				fmt.Println(err)
			}
			return err
		}, func(e error) error {
			response, err = fallbackProds()
			return err
		})

		if err != nil {
			ginCtx.JSON(500, gin.H{"status": err.Error()})
		} else {
			ginCtx.JSON(http.StatusOK, gin.H{"data": response.Data})
		}*/
	}

}
