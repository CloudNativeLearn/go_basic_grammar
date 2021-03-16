package main

import (
	"fmt"
	"strconv"
)

// 给定一个数字，我们按照如下规则把它翻译为字符串：0 翻译成 “a” ，1 翻译成 “b”，……，11 翻译成 “l”，……，25 翻译成 “z”。一个数字可能有多个翻译。请编程实现一个函数，用来计算一个数字有多少种不同的翻译方法。
//
// 
//
//示例 1:
//
//输入: 12258
//输出: 5
//解释: 12258有5种不同的翻译，分别是"bccfi", "bwfi", "bczi", "mcfi"和"mzi"


func main() {
	fmt.Println(int('1'))
	s := strconv.Itoa(123)
	fmt.Println(s[:len(s)-2])
	fmt.Println(UseStark(12258))
}

func translateNum(num int) int {
return UseStark(num)
}

// 使用递归
func UseStark(num int) int {
	// 如果小于10
	// 如果大于10 小于26
	// 其他
	if num < 10 {
		return 1
	} else if num >= 10 && num < 26 {
		return 2
	}
	str:= strconv.Itoa(num)
	db := make([]int,len(str))
	db[0] = 1 //边界
	if str[:2]>="10"&&str[:2]<="25"{    //边界
		db[1]=2
	}else{
		db[1]=1
	}

	for i:=2;i<len(str);i++{
		newnum:=str[i-1:i+1]
		if newnum>="10" && newnum<="25"{
			db[i]=db[i-1]+db[i-2]     //递归关系式
		}else{
			db[i]=db[i-1]
		}
	}
	return db[len(str)-1]

}