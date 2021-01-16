package route

import (
	"github.com/gin-gonic/gin"
	"../ws"
)

func Init()  {
	router := gin.Default()
	router.GET("/ping",ws.Ping)
	router.Run(":3001")
}
