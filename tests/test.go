package main

import (
	"context"
	"flag"
	"fmt"
	"golang.org/x/time/rate"
	"log"
	"net/http"

	"time"
)

/**
go run test.go -h
go run test.go -name abc
*/

func flagtest() {
	name := flag.String("name", "", "服务名")
	flag.Parse()

	fmt.Println(*name)
}

func ratetest() {
	r := rate.NewLimiter(1, 5)
	ctx := context.Background()
	for {
		err := r.WaitN(ctx, 2) //block
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(time.Now().Format("2006-01-02 15:05:33"))
		time.Sleep(time.Second)
	}
}

func ratetest2() {
	r := rate.NewLimiter(1, 5)

	for {
		if r.AllowN(time.Now(), 2) {
			fmt.Println(time.Now().Format("2006-01-02 15:05:33"))
		} else {
			fmt.Println("too many request")
		}
		time.Sleep(time.Second)
	}
}

var r = rate.NewLimiter(1, 5)

func MyLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if r.Allow() {
			next.ServeHTTP(writer, request)
		} else {
			http.Error(writer, "too many request", http.StatusTooManyRequests)
		}
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("OK!"))
	})
	http.ListenAndServe(":8888", MyLimit(mux))
}
