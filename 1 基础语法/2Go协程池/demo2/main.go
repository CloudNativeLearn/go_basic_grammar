package main

import "fmt"

func go_word(c chan int)  {
	fmt.Println("线程")
	// 如果不能从管道中获得对应的元素 会一直阻塞在这里
	num := <-c
	fmt.Println("从主线程中得到数据",num)
}

func main() {
	c:= make(chan int)
	go go_word(c)
	c <-2
	fmt.Println("im main")
}
