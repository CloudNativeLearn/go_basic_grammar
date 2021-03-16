package main

import "fmt"

// 给你一个字符串 s ，它只包含三种字符 a, b 和 c 。
//
//请你返回 a，b 和 c 都 至少 出现过一次的子字符串数目。
//

//示例 1：
//输入：s = "ab ca bc"
//输出：10
//解释：包含 a，b 和 c 各至少一次的子字符串为 "abc", "abca", "abcab", "abcabc", "bca", "bcab", "bcabc", "cab", "cabc" 和 "abc" (相同字符串算多次)。

//示例 2：
//输入：s = "aaacb"
//输出：3
//解释：包含 a，b 和 c 各至少一次的子字符串为 "aaacb", "aacb" 和 "acb" 。

//示例 3：
//输入：s = "abc"
//输出：1

func main() {
 s := "123"
 fmt.Println(string(s[1]))
}

func numberOfSubstrings(s string) int {
	begin := 0
	end := 2
	ThreSult := map[string]int{}
	ThreSult["a"] = 0
	ThreSult["b"] = 0
	ThreSult["c"] = 0
	ThreSult[string(s[0])] = ThreSult[string(s[0])]+1
	ThreSult[string(s[1])] = ThreSult[string(s[1])]+1
	ThreSult[string(s[2])] = ThreSult[string(s[2])]+1
	result := 0
	for end !=begin{
		// 如果已经满足 abc都有 则左边指针又移
		// 如果不满足   abc都有 则右边指针又移动
		if ThreSult["a"]==0|| ThreSult["b"] ==0 || ThreSult["c"] ==0{
			if end==len(s)-1{
				return result
			}
			ThreSult[string(s[end+1])] = ThreSult[string(s[end+1])] +1
			end = end +1
		}else {
			// 保存结果
			// 左指针右移
			// 结果减一
			result = result + len(s) -end
			ThreSult[string(s[begin])] = ThreSult[string(s[begin])] -1
			begin = begin+1

		}

	}

	return result
}
