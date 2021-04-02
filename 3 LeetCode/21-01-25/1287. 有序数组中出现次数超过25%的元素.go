package main

import "fmt"

// 给你一个非递减的有序整数数组，已知这个数组中恰好有一个整数，它的出现次数超过数组元素总数的 25%。
//
//请你找到并返回这个整数
//示例：
//输入：arr = [阿里笔试,2,2,6,6,6,6,7,10]
//输出：6


func main() {
	list := []int{1,2,2,6,6,6,6,7,10}
	fmt.Println(findSpecialInteger(list))
}

func findSpecialInteger(arr []int) int {
	theMap := map[int]int{}
	for _,i := range arr{
		theMap[i] = theMap[i] +1
	}
	theMaxNumber := -1
	theMax := -1
	for k,v := range theMap{
		if v >theMaxNumber {
			theMaxNumber = v
			theMax = k
		}
	}
	return theMax
}
