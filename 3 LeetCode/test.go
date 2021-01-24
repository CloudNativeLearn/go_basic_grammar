package main

import "fmt"

var TheParent = map[int]int{}
func main() {

	fmt.Println(TheParent)
	TheParent[0] = 1
	fmt.Println(TheParent[0])
	fmt.Println(TheParent)
	fmt.Println(TheParent[9])


}
