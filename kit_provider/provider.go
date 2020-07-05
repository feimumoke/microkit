package main

import (
	"flag"
	"fmt"
	kitlog "github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	mymux "github.com/gorilla/mux"
	"golang.org/x/time/rate"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	. "zhuhui.com/microkit/kit_provider/ps"
	"zhuhui.com/microkit/kit_provider/util"
)

func main() {

	name := flag.String("name", "userservice", "服务名称")
	port := flag.Int("p", 8080, "服务端口")
	flag.Parse()
	if *name == "" || *port == 0 {
		log.Fatal("please input service name and port")
	}
	util.SetUserServiceNameAndPort(*name, *port)
	var logger kitlog.Logger
	{
		logger = kitlog.NewLogfmtLogger(os.Stdout)
		logger = kitlog.WithPrefix(logger, "mykit", "1.0")
		logger = kitlog.With(logger, "time", kitlog.DefaultTimestampUTC)
		logger = kitlog.With(logger, "caller", kitlog.DefaultCaller)
	}

	userService := UserService{}
	limit := rate.NewLimiter(1, 5)

	//装饰器模式
	endp := RateLimit(limit)(UserServiceLogMiddleware(logger)(CheckTokenMiddleware()(GenUserEndponit(userService))))

	options := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(MyErrorEncoder),
	}
	handler := kithttp.NewServer(endp, DecodeUserRequest, EncodeUserResponse, options...)

	accessService := &AccessService{}
	accessEndponit := AccessEndponit(accessService)
	accessHandler := kithttp.NewServer(accessEndponit, DecodeAccessRequest, EncodeAccessResponse, options...)

	router := mymux.NewRouter()
	{
		//router.Handle(`/user/{uid:\d+}`, handler)
		router.Methods("POST").Path("/access-token").Handler(accessHandler)
		router.Methods("GET", "DELETE").Path(`/user/{uid:\d+}`).Handler(handler)
		router.Methods("GET").Path("/health").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			writer.Header().Set("Content-type", "application/json")
			writer.Write([]byte(`{"status":"ok"}`))
		})
	}

	errChan := make(chan error)

	go func() {
		util.RegService()
		err := http.ListenAndServe(":"+strconv.Itoa(*port), router)
		if err != nil {
			log.Println(err)
			errChan <- err
		}
	}()

	go func() {
		sig := make(chan os.Signal)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-sig)
	}()

	getErr := <-errChan
	util.Unregservice()
	log.Println(getErr)

}
