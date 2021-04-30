package main

// 面试题 16.26. 计算器
//给定一个包含正整数、加(+)、减(-)、乘(*)、除(/)的算数表达式(括号除外)，计算其结果。
//
//表达式仅包含非负整数，+， - ，*，/ 四种运算符和空格  。 整数除法仅保留整数部分。
//
//示例 1:
//
//输入: "3+2*2"
//输出: 7
//示例 2:
//
//输入: " 3/2 "
//输出: 1
//示例 3:
//
//输入: " 3+5 / 2 "
//输出: 5
//说明：
//
//你可以假设所给定的表达式都是有效的。
//请不要使用内置的库函数 eval。

func calculate(s string) int {
	stark := []int{}
	Fuhao := '+'
	num := 0
	for i, v := range s {
		isNumber := v <= '9' && v >= '0'
		if isNumber{
			num = num*10 + int(v - '0')
		}

		if !isNumber && v!=' ' || i==len(s)-1{
			switch Fuhao {
			case '+':
				stark = append(stark, num)
			case '-':
				stark = append(stark, -num)
			case '*':
				stark[len(stark)-1] = stark[len(stark)-1]*num
			case '/':
				stark[len(stark)-1] = stark[len(stark)-1]/num
			}

			Fuhao = v
			num = 0
		}



	}

	ans := 0
	for _, v := range stark {
		ans += v
	}
	return ans
}
