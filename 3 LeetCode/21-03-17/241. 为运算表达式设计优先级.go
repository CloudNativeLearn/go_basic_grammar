package main

import (
	"fmt"
	"strconv"
)

// 给定一个含有数字和运算符的字符串，为表达式添加括号，改变其运算优先级以求出不同的结果。你需要给出所有可能的组合的结果。有效的运算符号包含 +, - 以及 * 。
//
//示例 1:
//
//输入: "2-1-1"
//输出: [0, 2]
//解释:
//((2-1)-1) = 0
//(2-(1-1)) = 2
//示例 2:
//
//输入: "2*3-4*5"
//输出: [-34, -14, -10, -10, 10]
//解释:
//(2*(3-(4*5))) = -34
//((2*3)-(4*5)) = -14
//((2*(3-4))*5) = -10
//(2*((3-4)*5)) = -10
//(((2*3)-4)*5) = 10

func main() {
	s := "1+3"
	fmt.Println(s[0] - 48)
	fmt.Println(string(s[1]) == "+")
	fmt.Println(s[2] - 48)

	for _, v := range s {
		fmt.Println(v)
	}
}

func diffWaysToCompute(expression string) []int {
	k, err := strconv.Atoi(expression)
	if err == nil {
		// 没有返回错误 表示字符串为纯数字
		return []int{k}
	}

	result := []int{}
	// 循环便利字符串
	for k, v := range expression {
		// 如果是加减乘除
		if string(v) == "+" || string(v) == "*" || string(v) == "-" {
			left := diffWaysToCompute(expression[:k])
			right := diffWaysToCompute(expression[k+1:])
			for _,l := range left{
				for _,r := range right{
					switch string(v) {
					case "+":
						result = append(result, l+r)
					case "-":
						result= append(result, l-r)
					case "*":
						result = append(result, l*r)


					}
				}
			}
		}
	}
	return result
}
