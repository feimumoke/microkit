package main

import (
	"errors"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"k8s.io/apimachinery/pkg/util/rand"
	"sync"
	"time"
)

type Product struct {
	ID    int
	Title string
	Price int
}

func getProduct() (Product, error) {
	r := rand.Intn(10)
	if r < 5 {
		time.Sleep(time.Second * 3)
	}
	return Product{
		ID:    101,
		Title: "skill of go",
		Price: 60,
	}, nil
}

func RecProduct() (Product, error) {
	return Product{
		ID:    0,
		Title: "recommand book",
		Price: 30,
	}, nil

}

func syncHys() {
	configA := hystrix.CommandConfig{Timeout: 2000,}
	hystrix.ConfigureCommand("get_prod", configA)
	for {
		err := hystrix.Do("get_prod", func() error { //Do 同步 Go 异步
			p, _ := getProduct()
			fmt.Println(p)
			return nil
		}, func(e error) error {
			fmt.Println(RecProduct())
			return errors.New("my time out")
		})
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(time.Second)
	}
}
func main2() {
	rand.Seed(time.Now().UnixNano())
	configA := hystrix.CommandConfig{
		Timeout:                2000, //超时时间
		MaxConcurrentRequests:  2,    //最大并发数
		RequestVolumeThreshold: 5,    //熔断器请求阈值，有5个请求才进行错误百分比计算
		ErrorPercentThreshold:  99,
		SleepWindow:            10}   //尝试去请求
	hystrix.ConfigureCommand("get_prod", configA)
	c, _, _ := hystrix.GetCircuit("get_prod")
	resultChan := make(chan Product, 1)
	wg := sync.WaitGroup{}
	for i := 0; i < 4; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				wg.Add(1)
				defer wg.Done()
				errs := hystrix.Go("get_prod", func() error { //Do 同步 Go 异步
					p, _ := getProduct()
					resultChan <- p
					return nil
				}, func(e error) error {
					p, _ := RecProduct()
					resultChan <- p
					return errors.New("my time out")
				})

				select {
				case prod := <-resultChan:
					fmt.Println(prod)
				case err := <-errs:
					fmt.Println("err:", err, "===")
				}
				fmt.Println(c.IsOpen())
				time.Sleep(time.Second)
			}
		}()
	}

	wg.Wait()

	time.Sleep(time.Hour)

}
