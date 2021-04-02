package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

func main() {
	userList = make([]map[string]string,10,10)
	listChatHistory = make([]map[string]string,100,100)
	router := gin.Default()
	router.GET("/ping",Ping)
	router.Run(":3001")
}
var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 在线列表  {  name,time }
var userList []map[string]string
// 所有人的聊天信息  {  name message time towho}
var listChatHistory []map[string]string
// 返回的消息
var ReturnMessage = &Themessage{}

type Themessage struct {
	Name string `json:"name"`
	AllInUser []map[string]string `json:"all_in_user"`
	ListChatHistory  []map[string]string `json:"list_chat_history"`
	Message string
}


func Ping(c *gin.Context)  {
	// 升级get请求为websocket协议
	ws,err := upGrader.Upgrade(c.Writer,c.Request,nil)
	if err != nil{
		return
	}
	defer ws.Close()

	var Username string
	defer loginout(Username)

	for{
		// 读取ws中的数据
		_,message,err := ws.ReadMessage()
		if err != nil{
			// 客户端关闭连接时也会进入
			fmt.Println(err)
			break
		}


		var d map[string]interface{}
		// 将字符串反解析为字典
		fmt.Println(string(message))

		err = json.Unmarshal(message, &d)
		if err != nil{
			fmt.Println("传输失败")
		}
		fmt.Println(d)  // map[id:阿里笔试 name:wxnacy]

		switch d["sendType"] {
		case "login":
			//
			// { name :"","sendType":"login"}
			//
			Username = d["name"].(string)
			_,m := login(d["name"].(string))
			err = ws.WriteJSON(m)
			if err != nil{
				break
			}
		case "sendmessage":
			Username = d["name"].(string)
			// {name :"","sendType":"sendmessage",userlist:["",""]}
			m :=SendMessage(d["userlist"].([]string),d["message"].(string),d["name"].(string))
			err = ws.WriteJSON(m)
			if err != nil{
				break
			}
		}



		//
		//// 如果客户端发送ping就返回pong 否则数据原封不动返回客户端
		//if string(message) == "ping"{
		//	message = []byte("pong")
		//}
		//// 写入ws数据 二进制返回
		//err = ws.WriteMessage(mt,message)
		//// 返回json字符串，借助gin的ginH实现
		//// v := gin.H{"message":msg}
		//// err = ws.WriteJSON(v)
		//fmt.Println(message)
		//v := gin.H{"message":message,"数据返回":"数据返回"}
		//err = ws.WriteJSON(v)
		//if err != nil{
		//	break
		//}


	}




}


// 登录  { name : "" }
func login(name string) (bool ,*Themessage) {
	for _, value := range userList{
		if value["name"] == name{
			return false,nil
		}
	}
	user := map[string]string{
		"name":name,
		"time":time.Now().String(),
	}
	userList = append(userList, user)
	message := &Themessage{
		Name: name,
		AllInUser: userList,
		ListChatHistory: listChatHistory,
		Message: "登录成功",
	}
	fmt.Println("登录成功返回数据")
	fmt.Println(message)

	return true,message
}

// 退出登录的时候执行
func loginout(name string)  {
	for i, value := range userList{
		if value["name"] == name{
			userList = append(userList[:i],userList[i+1:]... )
		}
	}
}

// 给某个人发送消息
func SendMessage(userlist22 []string,message string,name string) *Themessage {
	for _,value := range userlist22{
		message := map[string]string{
			"name":name,
			"message":message,
			"time":time.Now().String(),
			"towho":value,
		}
		listChatHistory = append(listChatHistory,message)
	}
	returnmessage := &Themessage{
		Name: name,
		AllInUser: userList,
		ListChatHistory: listChatHistory,
		Message: "发送成功",
	}

	return returnmessage

}







//// 给某个人发送信息  { myname toid  message }
//func sendToOne(data map[string]interface{})  {
//
//	message := map[string]string{
//		"name" : data["myname"].(string),
//		"message":data["message"].(string),
//		"time": time.Now().String(),
//
//	}
//}




