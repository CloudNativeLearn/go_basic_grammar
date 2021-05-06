package main

import "fmt"

// 闭包的应用

func a() func() int {
	var x int
	return func() int {
		x = x + 1
		return x * x
	}
}

func main() {
	aa := a()
	fmt.Println(aa())
	fmt.Println(aa())
	fmt.Println(aa())

	bb := func() func() int {
		var x int
		return func() int {
			x = x + 1
			return x
		}
	}()

	fmt.Println(bb())
	fmt.Println(bb())
	fmt.Println(bb())

}
