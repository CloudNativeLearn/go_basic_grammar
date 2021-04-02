package main

import (
	"bufio"
	"fmt"
	"net"
)

// 服务端

// 处理函数
func process(conn net.Conn)  {
	defer conn.Close() // 关闭链接
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n,err:= reader.Read(buf[:]) //读取数据
		if err != nil{
			fmt.Println("reader from client failde ,err:",err)
			break
		}

		recvStr := string(buf[:n])
		fmt.Println("收到client端发来的数据",recvStr)
		conn.Write([]byte(recvStr))
	}
}

func main() {
	listen,err := net.Listen("tcp","127.0.0.阿里笔试:20000")
	if err!= nil{
		fmt.Println("listen failed , err:",err)
		return
	}

	for{
		conn,err := listen.Accept()// 建立链接
		if err!= nil{
			fmt.Println("accept failed, err:",err)
			continue
		}
		go process(conn) //启动一个携程处理链接
	}
}