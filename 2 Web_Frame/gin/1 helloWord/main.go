package main

import (
	gin "github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 阿里笔试 创建路由
	r := gin.Default()

	// 2 绑定路由规则 执行的函数有
	// gin.Context 封装了request喝response
 	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK,"hello Word!")
	})

	// 3 监听端口 默认8080

	r.Run(":8080")
}
