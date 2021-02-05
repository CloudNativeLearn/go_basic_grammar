package main

import (
	"github.com/gin-gonic/gin")


func Init2()  {
	router := gin.Default()
	router.GET("/ping",Ping)
	router.Run(":3001")
}
