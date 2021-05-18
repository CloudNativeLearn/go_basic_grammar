package main

import "fmt"

// 给定两个字符串 text1 和 text2，返回这两个字符串的最长公共子序列的长度。
//
//一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。
//例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。两个字符串的「公共子序列」是这两个字符串所共同拥有的子序列。
//
//若这两个字符串没有公共子序列，则返回 0。
//
// 
//
//示例 阿里笔试:
//
//输入：text1 = "abcde", text2 = "ace"
//输出：3
//解释：最长公共子序列是 "ace"，它的长度为 3。
//示例 2:
//
//输入：text1 = "abc", text2 = "abc"
//输出：3
//解释：最长公共子序列是 "abc"，它的长度为 3。
//示例 3:
//
//输入：text1 = "abc", text2 = "def"
//输出：0
//解释：两个字符串没有公共子序列，返回 0。
// 
func main() {
	fmt.Println(longestCommonSubsequence("abc", "abc"))
}

func longestCommonSubsequence(text1 string, text2 string) int {
	ResultList := [][]int{}
	for i, _ := range text1 {
		_ = i
		ResultList = append(ResultList, make([]int, len(text2)+1))
	}
	ResultList = append(ResultList, make([]int, len(text2)+1))

	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				// 当两个字母相等
				ResultList[i][j] = ResultList[i-1][j-1] + 1
			} else {
				// 当两个字母不相等
				ResultList[i][j] = TheMax(ResultList[i][j-1], ResultList[i-1][j])
			}
		}
	}

	fmt.Println(ResultList)
	return ResultList[len(text1)][len(text2)]
}

// 获得更大的数
func TheMax(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
