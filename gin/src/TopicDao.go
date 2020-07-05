package src

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func MustLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, status := c.GetQuery("token"); !status {
			c.String(http.StatusUnauthorized, "do not hava token")
			c.Abort()
		} else {
			c.Next()
		}
	}
}

func GetTopicDetail(context *gin.Context) {
	//context.JSON(http.StatusOK,CreateTopic(01,"Topic title"))
	//context.String(http.StatusOK, "Get topic Id : %s", context.Param("topic_id"))
	tid := context.Param("topic_id")
	topics := Topics{}
	DBHelper.Find(&topics, tid)

	context.Set("dbResult",topics)

	//
	//conn := RedisDefaultPool.Get()
	//defer conn.Close() // back to the pool, not closed
	//redisKey := "topic_" + tid
	//res, err := redis.Bytes(conn.Do("get", redisKey))
	//if err != nil {
	//	DBHelper.Find(&topics, tid)
	//	retData, _ := ffjson.Marshal(topics)
	//	if topics.TopicId == 0 {
	//		//no data in db, avoid cache shot
	//		conn.Do("setex", redisKey, 10, retData)
	//	} else {
	//		// normal
	//		conn.Do("setex", redisKey, 60, retData)
	//	}
	//	log.Println("read from mysql")
	//} else {
	//	ffjson.Unmarshal(res, &topics)
	//	fmt.Println("read from redis")
	//}

	//context.JSON(http.StatusOK, topics)
}

func NewTopic(c *gin.Context) {
	topic := Topics{}
	err := c.BindJSON(&topic)
	if err != nil {
		c.String(http.StatusBadRequest, "Param error: %s", err.Error())
	} else {
		c.JSON(http.StatusOK, topic)
	}
	//c.String(http.StatusOK,"New Topic")
}

func NewTopics(c *gin.Context) {
	topics := TopicBox{}
	err := c.BindJSON(&topics)
	if err != nil {
		c.String(http.StatusBadRequest, "Param error: %s", err.Error())
	} else {
		c.JSON(http.StatusOK, topics)
	}
	//c.String(http.StatusOK,"New Topic")
}

func DelTopic(c *gin.Context) {

	c.String(http.StatusOK, "Del Topic")
}
func GetTopicList(c *gin.Context) {
	//if c.Query("username") == "" {
	//	c.String(http.StatusOK, "Get topic list")
	//} else {
	//	c.String(http.StatusOK, "Get topic of : %s", c.Query("username"))
	//}
	query := TopicQuery{}
	err := c.BindQuery(&query)
	if err != nil {
		c.String(http.StatusBadRequest, "Param error: %s", err.Error())
	} else {
		c.JSON(http.StatusOK, query)
	}
}
