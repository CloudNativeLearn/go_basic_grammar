package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"reflect"
	"encoding/json"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// ping    websocket请求ping 返回pong
func Ping(c *gin.Context)  {
	// 升级get请求为websocket协议
	ws,err := upGrader.Upgrade(c.Writer,c.Request,nil)
	if err != nil{
		return
	}
	defer ws.Close()

	for{
		// 读取ws中的数据
		mt,message,err := ws.ReadMessage()
		if err != nil{
			// 客户端关闭连接时也会进入
			fmt.Println(err)
			break
		}
		fmt.Println("开始从前端读取数据")
		fmt.Println(mt)
		fmt.Println(message)
		fmt.Println(reflect.TypeOf(message))

		fmt.Println(string(message))

		var d map[string]interface{}
		// 将字符串反解析为字典
		err = json.Unmarshal(message, &d)
		if err != nil{
			fmt.Println("传输失败")
		}
		fmt.Println(d)  // map[id:阿里笔试 name:wxnacy]



		fmt.Printf("读取结束\n")

		// 如果客户端发送ping就返回pong 否则数据原封不动返回客户端
		if string(message) == "ping"{
			message = []byte("pong")
		}
		// 写入ws数据 二进制返回
		err = ws.WriteMessage(mt,message)
		// 返回json字符串，借助gin的ginH实现
		// v := gin.H{"message":msg}
		// err = ws.WriteJSON(v)
		fmt.Println(message)
		v := gin.H{"message":message,"数据返回":"数据返回"}
		err = ws.WriteJSON(v)
		if err != nil{
			break
		}
	}

}