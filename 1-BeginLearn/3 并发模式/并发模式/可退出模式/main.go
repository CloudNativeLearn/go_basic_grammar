package main

import (
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	quit := make(chan bool)
	for i:=0;i<2;i++{
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			task := func() {
				println(id,time.Now().Nanosecond())
				time.Sleep(time.Second)
			}
			for{
				select {
				case <-quit:  // closed channel 不会阻塞，因此可以用于推出通知
					return
				default:
					task()
				}
			}
		}(i)

	}
	time.Sleep(time.Second*5) // 让测试goroutine运行一会
	close(quit) //发出退出通知
	wg.Wait()
}


