package main

import (
	"fmt"
	"time"
)

func go_word(name string) {
	for i := 0; i < 10; i++ {
		fmt.Println("我是go的一个携程", name)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("携程执行完毕", name)
}

func main() {
   //
	go go_word("小黑")
	go go_word("小白")
	for i:=0;i<10;i++{
		fmt.Println("main主函数")
		time.Sleep(1*time.Second)
	}
}
