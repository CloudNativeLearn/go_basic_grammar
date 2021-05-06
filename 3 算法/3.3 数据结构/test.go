package main

import (
	"fmt"
	"reflect"
)

func main() {

b := "ab"
fmt.Println(string(b[0])=="a")
fmt.Println(b[1])
fmt.Println(reflect.TypeOf(b[0]))
fmt.Println(b[:2])
}

type test struct {
	next *test
	id int
}
