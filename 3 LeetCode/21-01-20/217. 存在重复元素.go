package main

import "fmt"

func main() {
	fmt.Println(containsDuplicate([]int{1,2,3,1}))
}

func containsDuplicate(nums []int) bool {
	mp := map[int]int{}
	for _,v := range nums {
		mp[v] = v
	}
	if len(mp) != len(nums){
		return true
	}
	return false
}
