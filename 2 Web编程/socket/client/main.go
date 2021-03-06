package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// 阿里笔试 与服务器建立连接
	conn,err := net.Dial("tcp","127.0.0.阿里笔试:20000")
	if err != nil{
		fmt.Printf("dial failed,err:%v\n",err)
		return
	}
	// 2 利用连接进行数据的发送和接收
	input := bufio.NewReader(os.Stdin)
	for {
		s,_ := input.ReadString('\n')
		s = strings.TrimSpace(s)
		if strings.ToUpper(s) == "Q"{
			return
		}
		// 给服务器发送消息
		_,err := conn.Write([]byte(s))
		if err != nil{
			fmt.Printf("send failed,err:%v\n",err)
			return
		}


		// 从服务端接收回复的消息
		var buf [1024] byte
		n,err := conn.Read(buf[:])
		if err != nil{
			fmt.Printf("reader failed ,err:%v\n",err)
			return
		}
		fmt.Println("收到服务端回复：",string(buf[:n]))


	}
}
