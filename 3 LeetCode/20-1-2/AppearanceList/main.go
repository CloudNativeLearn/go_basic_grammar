package main

import (
	"fmt"
	"strconv"
)

// 给定一个正整数 n（1 ≤ n ≤ 30），输出外观数列的第 n 项。
//
//注意：整数序列中的每一项将表示为一个字符串。
//
//「外观数列」是一个整数序列，从数字 1 开始，序列中的每一项都是对前一项的描述。前五项如下：
//
//1.     1
//2.     11
//3.     21
//4.     1211
//5.     111221
//第一项是数字 1
//
//描述前一项，这个数是 1 即 “一个 1 ”，记作 11
//
//描述前一项，这个数是 11 即 “两个 1 ” ，记作 21
//
//描述前一项，这个数是 21 即 “一个 2 一个 1 ” ，记作 1211
//
//描述前一项，这个数是 1211 即 “一个 1 一个 2 两个 1 ” ，记作 111221
//
// 
//
//示例 1:
//
//输入: 1
//输出: "1"
//解释：这是一个基本样例。
//示例 2:
//
//输入: 4
//输出: "1211"
//解释：当 n = 3 时，序列是 "21"，其中我们有 "2" 和 "1" 两组，"2" 可以读作 "12"，也就是出现频次 = 1 而 值 = 2；类似 "1" 可以读作 "11"。所以答案是 "12" 和 "11" 组合在一起，也就是 "1211"。


// 执行用时：
//4 ms
//, 在所有 Go 提交中击败了
//47.94%
//的用户
//内存消耗：
//6.4 MB
//, 在所有 Go 提交中击败了
//29.03%
//的用户
func countAndSay(n int) string {
	return "1"
}

var Theresult string

func StartIt(n int, Theresult *string) string {
	if n == 1 {
		*Theresult = "1"
		return *Theresult
	} else {
		// 不等于一 递归调用上一个
		TheOldResult := StartIt(n-1, Theresult)
		*Theresult = ""
		// 循环旧数据进行判断
		// 需要两个字段   Theword 记录当前数字   ThwordNumber 记录出现次数
		var Theword string
		var TheWordNumber int
		// 开始循环旧数字
		for i := 0; i < len(TheOldResult); i++ {

			if i == 0 {
				Theword = fmt.Sprintf("%c", TheOldResult[i])
				TheWordNumber = 1
				if len(TheOldResult) == 1 {
					*Theresult = *Theresult + strconv.Itoa(TheWordNumber)+Theword
					continue
				}
				continue
			}

			if Theword == fmt.Sprintf("%c", TheOldResult[i]) {
				// 两个数字一样
				TheWordNumber = TheWordNumber+1
				if i == len(TheOldResult)-1{
					*Theresult = *Theresult+strconv.Itoa(TheWordNumber)+Theword
					continue
				}

			} else {
				// 两个数字不一样
				//先加入
				*Theresult = *Theresult+strconv.Itoa(TheWordNumber)+Theword
				// 后更改
				Theword = fmt.Sprintf("%c", TheOldResult[i])
				TheWordNumber = 1
				//在判断是否最后一个
				if i == len(TheOldResult)-1{
					*Theresult = *Theresult+strconv.Itoa(TheWordNumber)+Theword
					continue
				}


			}

		}


		return *Theresult

	}

}



func main() {
	StartIt(5, &Theresult)




}
