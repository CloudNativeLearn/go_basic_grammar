package main

import (
	"runtime"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i:=0;i<6;i++{
			println(i)
			if i==3{
				runtime.Gosched()  // 礼让线程 关闭线程
			}
		}
	}()

	go func() {
		defer wg.Done()
		println("Hello,word!!")
	}()

	wg.Wait()
}
