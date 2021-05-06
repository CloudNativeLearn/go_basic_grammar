package main

import (
	"fmt"
	"os"
)
// 主要特征  有多个管道的时候可以随机选择管道存入数据
func main() {
	a,b := make(chan int,3),make(chan int)
	go func() {
		v,ok,s := 0,false,""
		for{
			select {
				case v,ok =<-a:
					s="a"
				case v,ok=<-b:
					s="b"
			}
			if ok{
				fmt.Println(s,v)
			}else {
				os.Exit(0)
			}
		}
	}()

	for i:=0;i<5;i++{
		// 随机选择可用channel，发送数据
		select {
		case a<-i:
		case b<-i:
		}
	}
	close(a)
	select {
	// 没有可用channel,阻塞main goroutine
	}
}
