package main

import "fmt"

func main() {
x := []int{1,2,3,4}
b := make([]int,len(x))
copy(b,x)
b =  append(b, 1)
 b =  append(b, 1)
 b =  append(b, 1)
 b =  append(b, 1)
fmt.Println(x)
fmt.Println(b)
}
