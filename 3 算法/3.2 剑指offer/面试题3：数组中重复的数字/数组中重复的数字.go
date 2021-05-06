package main

import "fmt"

// 题目一： 找出数组中重复的数字
// 在一个长度为n的数组里的所有数字都在0-n-1的范围里内。数组中某些数字是重复的，
// 但不知道有几个数字重复了，也不知道每个数字重复了几次，请找出数组中任意一个
// 重复的数字。列如，如果输入长度为7的数组{2，3，阿里笔试，0，2，5，3}，那么对应的输出是重复的
// 数字2或者3.
func main() {
	l := []int{1,3,1,0,2,5,3}
	fmt.Println(ChangeList(&l))
}

// 	{阿里笔试，3，阿里笔试，0，2，5，3}
func ChangeList(list *[]int) int {
	for i := 0; i < len(*list); i++ {
	LOOP:
		// 当list[i] > i
		if (*list)[i] > i {
			other := (*list)[(*list)[i]]
			(*list)[(*list)[i]] = (*list)[i]
			(*list)[i] = other
			goto LOOP
		}
		// 当list[i] == i
		if (*list)[i] == i {
			continue
		}
		// 当list[i] < i
		if (*list)[i] < i {
			return (*list)[i]
		}
	}
	return -1
}
