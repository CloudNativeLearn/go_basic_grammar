package main

import (
	"fmt"
	"sync"
)
// 主要特点 使用这种模式之后
// 线程不会轮询执行会进行阻塞
// 完成当前线程任务之后再进行
// 其余线程的执行
func main() {
	wg := sync.WaitGroup{}
	wg.Add(3)

	sem := make(chan int,3)
	for i:=0;i<3;i++{
		go func(id int) {
			defer wg.Done()
			sem <-1
			for x:=0;x<3;x++{
				//fmt.Println(id,x)
				fmt.Println("线程id",id,"数据为",x)
			}
			<- sem
		}(i)
	}

	wg.Wait()


}
