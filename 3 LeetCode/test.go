package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := []int{1, 2, 3}
	fmt.Println(a[0 : len(a)-1])
	b := " */"
	for _, v := range b {
		fmt.Println(v)
		fmt.Println(reflect.TypeOf(v))
	}

}
