package main

import "fmt"

func main() {
a := []int{1,2,3}
fmt.Println(a[:1],a[1:],a[:len(a)-1])

}
