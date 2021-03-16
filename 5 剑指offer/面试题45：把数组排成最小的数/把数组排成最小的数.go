package main

import (
	"fmt"
	"strconv"
)

// 输入一个非负整数数组，把数组里所有数字拼接起来排成一个数，打印能拼接出的所有数字中最小的一个。
//
// 
//
//示例 1:
//
//输入: [10,2]
//输出: "102"
//示例 2:
//
//输入: [3,30,34,5,9]
//输出: "3033459"
//
func main() {
nums:= []int{3,30,34,5,9}
fmt.Println(minNumber(nums))
}



func minNumber(nums []int) string {
	result:=hill(nums)
	return  result
}


// 冒泡排序版本
func maopao(nums []int)string  {
	for i:=0;i<len(nums);i++ {
		for j:=i+1;j<len(nums);j++ {
			// 先直接判断两个

			// 讲两个数字变为字符串拼接比较大小
			a := strconv.Itoa(nums[i])
			b := strconv.Itoa(nums[j])
			result1,_ :=strconv.Atoi(a+b)
			result2,_:=strconv.Atoi(b+a);
			if result1>result2{
				// 如果大于则交换两个数字
				temp := nums[j]
				nums[j] = nums[i]
				nums[i] = temp
			}
		}
	}

	return ChangeListToString(nums)
}


// 希尔排序版本
func hill(nums []int)string  {
	numbers:= len(nums)
	for numbers!=1 {
		numbers = numbers/2
		for i:=numbers;i<len(nums);i++ {
			for j:=i-numbers;j>=0;j=j-numbers{
				a := strconv.Itoa(nums[j])
				b := strconv.Itoa(nums[j+numbers])
				result1,_ :=strconv.Atoi(a+b)
				result2,_:=strconv.Atoi(b+a);
				if result1 > result2{
					temp := nums[j]
					nums[j] = nums[j+numbers]
					nums[j+numbers] = temp
				}
			}
		}
	}

	return ChangeListToString(nums)
}

// 讲数组转换为字符串输出
func ChangeListToString(nums []int) (result string) {

	for _,v := range nums{
		result = result+strconv.Itoa(v)
	}
	return  result
}

