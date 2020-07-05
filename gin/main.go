package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"zhuhui.com/microkit/gin/src"
)

// set GO111MODULE=on

func v1(router gin.Engine) {
	router.GET("/v1/topics", func(context *gin.Context) {
		if context.Query("username") == "" {
			context.String(http.StatusOK, "Get topic list")
		} else {
			context.String(http.StatusOK, "Get topic of : %s", context.Query("username"))
		}

	})
	router.GET("/v1/topics/:topic_id", func(context *gin.Context) {
		context.String(http.StatusOK, "Get topic Id : %s", context.Param("topic_id"))

	})
}

func db() {
	db, _ := gorm.Open("mysql", "root:tiger@/gin?charset=utf8mb4&parseTime=True&loc=Local")
	topics := src.Topics{
		TopicTitle:      "别人的失败就是我的快乐啦",
		TopicShortTitle: "快乐男孩",
		UserIP:          "127.0.0.2",
		TopicScore:      0,
		TopicUrl:        "hahaha",
		TopicDate:       time.Now(),
	}
	fmt.Println(db.Create(&topics).RowsAffected)

	tc := src.TopicClass{}
	db.Table("topic_class").First(&tc, 2)
	fmt.Println(tc)
	rows, _ := db.Raw("select topic_id, topic_title from topics").Rows()
	for rows.Next() {
		var t_id int
		var t_title string
		rows.Scan(&t_id, &t_title)
		fmt.Println(t_id, t_title)
	}
	defer db.Close()
}

func os_signal() {
	count := 0
	go func() {
		for {
			fmt.Println("Execute,", count)
			count++
			time.Sleep(time.Second)
		}
	}()

	c := make(chan os.Signal)

	go func() {
		ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
		select {
		case <-ctx.Done():
			fmt.Println("Done")
			c <- os.Interrupt
		}
	}()

	signal.Notify(c)
	s := <-c
	fmt.Println(s)
}

func main2() {
	conn := src.RedisDefaultPool.Get()
	res, _ := redis.String(conn.Do("get", "a"))
	fmt.Println("get:", res)
}
func main() {
	//db()
	//os_signal()
	router := gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("topicurl", src.TopicUrl)
		v.RegisterValidation("topics", src.TopicsValidate)
	}

	v1 := router.Group("/v1/topics")
	{
		v1.GET("", src.GetTopicList)
		v1.GET("/:topic_id", src.CacheDecorator(src.GetTopicDetail, "topic_id", "topic_%s", src.Topics{}))
		v1.Use(src.MustLogin())
		{
			v1.POST("", src.NewTopic)
			v1.DELETE("/:topic_id", src.DelTopic)
		}
	}

	v2 := router.Group("/v1/mtopics")
	{
		v2.Use(src.MustLogin())
		{
			v2.POST("", src.NewTopics)
		}
	}

	//router.Run()
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("start server failed")
		}
	}()

	go func() {
		src.InitDB()
	}()
	src.ServerNotify()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	err := server.Shutdown(ctx)
	if err != nil {
		log.Fatal("Force shutdown", err)
	}
	log.Println("shut down graceful")
}
