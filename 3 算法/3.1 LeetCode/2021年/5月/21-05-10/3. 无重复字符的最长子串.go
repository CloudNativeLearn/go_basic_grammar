package main

import "fmt"

// 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
//
// 
//
//示例 1:
//
//输入: s = "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//示例 2:
//
//输入: s = "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//示例 3:
//
//输入: s = "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
//示例 4:
//
//输入: s = ""
//输出: 0
// 
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}

	TheMap := map[uint8]int{}
	i := 0
	j := 0
	Result := 1
	//TheMap[s[0]] = 1
	for {
		whether := WhetherRight(&TheMap)
		if whether && j < len(s) {
			Result = GetMax(Result, j-i)
			TheMap[s[j]] = TheMap[s[j]] + 1
			j = j + 1
		} else if i == len(s) && j == len(s) {
			return Result
		} else if whether && j >= len(s) {
			Result = GetMax(Result, j-i)
			TheMap[s[i]] = TheMap[s[i]] - 1
			i = i + 1
		} else {
			TheMap[s[i]] = TheMap[s[i]] - 1
			i = i + 1
		}

	}

}

func WhetherRight(Themap *map[uint8]int) bool {
	result := true
	for _, v := range *Themap {
		if v > 1 {
			result = false
		}
	}
	return result
}

func GetMax(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
