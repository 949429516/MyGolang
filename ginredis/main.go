package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func main() {
	r := gin.Default()
	v1 := r.Group("")
	{
		v1.GET("", bookshow)
	}
	r.Run(":8080")
}

func bookshow(c *gin.Context) {
	a := myredis()
	c.String(200, "有【%s】人访问了该网站", a)
}

func myredis() string {
	var rdb = redis.NewClient(&redis.Options{Addr: "localhost:6379", Password: "123456", DB: 1})
	_, err := rdb.Ping().Result()
	if err != nil {
		fmt.Println("ping 出错：", err)
	}
	rdb.Incr("age")
	res, err := rdb.Get("age").Result()
	if err != nil {
		fmt.Println("设置数据失败:", err)
	}
	return res
}
