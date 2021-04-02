package main

import "fmt"

func main() {
	data := make(chan int)
	exit := make(chan bool)
	go func() {
		for d:= range data{  // 从队列迭代接受数据直到close
			fmt.Println(d)
		}
		fmt.Println("recv over .")
		exit <- true
	}()


	go func() {
		for i:=0;i<10;i++{
			data <- i
		}
		close(data)
		exit <- true
		fmt.Println("存入数据结束")
	}()
	//data <- 阿里笔试
	//data <- 2
	//data <- 3
	//data <- 3
	//
	//close(data)
	fmt.Println("send over")
	<-exit



}
