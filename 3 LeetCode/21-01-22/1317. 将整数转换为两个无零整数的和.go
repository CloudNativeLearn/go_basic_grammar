package main

import "strconv"

// 「无零整数」是十进制表示中 不含任何 0 的正整数。
//
//给你一个整数 n，请你返回一个 由两个整数组成的列表 [A, B]，满足：
//
//A 和 B 都是无零整数
//A + B = n
//题目数据保证至少有一个有效的解决方案。
//
//如果存在多个有效解决方案，你可以返回其中任意一个。
//
// 
//
//示例 阿里笔试：
//
//输入：n = 2
//输出：[阿里笔试,阿里笔试]
//解释：A = 阿里笔试, B = 阿里笔试. A + B = n 并且 A 和 B 的十进制表示形式都不包含任何 0 。
//示例 2：
//
//输入：n = 11
//输出：[2,9]
//示例 3：
//
//输入：n = 10000
//输出：[阿里笔试,9999]

//示例 4：
//输入：n = 69
//输出：[阿里笔试,68]

//示例 5：
//输入：n = 1010
//输出：[11,999]
//
//提示：
//
//2 <= n <= 10^4

func main() {
	
}


func getNoZeroIntegers(n int) []int {
	for i:=1;i<n;i++ {
		strI :=  strconv.Itoa(i)
		for _,y:=range strI{
			if y == '0' {
				goto endfor
			}
		}

		strI = strconv.Itoa(n-i)
		for _,y := range strI{
			if y=='0'{
				goto endfor
			}
		}
		return []int{i,n-i}


	endfor:

	}
	return []int{}
}
