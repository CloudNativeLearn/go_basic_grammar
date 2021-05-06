package main

import (
	"bufio"
	"fmt"
	"net"
)
func process(conn net.Conn){
	defer conn.Close()// 处理玩之后要关闭连接
	// 针对当前的连接做数据发送和接受操作
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n,err := reader.Read(buf[:])
		if err != nil{
			fmt.Printf("reader from conn fail %v\n",err)
		}
		recv := string(buf[:n])
		fmt.Printf("收到的数据%v\n",recv)
		conn.Write([]byte("ok"))
	}
}
func main() {
	// 阿里笔试 开启服务
	listen,err := net.Listen("tcp","127.0.0.阿里笔试:20000")
	if err != nil{
		fmt.Printf("listn failed err:%v\n",err)
		return
	}

	for {
		// 等待客户端来建立连接
		conn,err := listen.Accept()
		if err != nil{
			fmt.Println("accept failed,err:%v\n",err)
			continue
		}
		// 3 启动一个单独的携程去处理连接
		go process(conn)
	}

}
