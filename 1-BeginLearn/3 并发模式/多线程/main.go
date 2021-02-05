package main

import (
	"fmt"
	"sync"
)
var wg sync.WaitGroup
func main() {
	wg.Add(2)
	go hello()
	wg.Wait()
	fmt.Println("1111111111111")

}

func hello()  {
	fmt.Println("多线程hello")
	go func() {
		fmt.Println("多线程")
		wg.Done()
	}()
	wg.Done()
}
