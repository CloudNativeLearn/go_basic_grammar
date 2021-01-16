package main

import (
	"fmt"
	"reflect"
)

func main() {
	var list *[]int
	mid := make([]int, 10)
	list = &mid
	mid[1] = 1



	//fmt.Println(list[:9])
	fmt.Println(list)
	fmt.Println(reflect.TypeOf(list))
	fmt.Println(reflect.TypeOf((*list)[0:1]))
	fmt.Println((*list)[1:2])
}
