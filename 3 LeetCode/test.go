package main

import "fmt"

func main() {
	ThMap := map[int]int{}
	ThMap[0] = 1
	ThMap[1] = 1
	for _,v := range ThMap{
		fmt.Println(v)
		ThMap[1] =2
	}

}