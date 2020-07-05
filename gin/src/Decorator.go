package src

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/pquerna/ffjson/ffjson"
	"net/http"
)

/**
缓存装饰器
 */
func CacheDecorator(h gin.HandlerFunc, param string, redisKeyPattern string, empty interface{}) gin.HandlerFunc {
	return func(context *gin.Context) {
		getId := context.Param(param)
		redisKey := fmt.Sprintf(redisKeyPattern, getId)
		conn := RedisDefaultPool.Get()
		defer conn.Close() // back to the pool, not closed
		res, err := redis.Bytes(conn.Do("get", redisKey))
		if err != nil {
			h(context)
			dbResult, exist := context.Get("dbResult")
			if !exist {
				dbResult = empty
			}
			retData, _ := ffjson.Marshal(dbResult)
			conn.Do("setex", redisKey, 60, retData)
			context.JSON(http.StatusOK, dbResult)
			//if topics.TopicId == 0 {
			//	//no data in db, avoid cache shot
			//	conn.Do("setex", redisKey, 10, retData)
			//} else {
			//	// normal
			//	conn.Do("setex", redisKey, 60, retData)
			//}

		} else {
			ffjson.Unmarshal(res, &empty)
			fmt.Println("read from redis")
			context.JSON(http.StatusOK, empty)
		}
	}
}
