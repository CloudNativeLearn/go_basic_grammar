package main

import "fmt"

//给你一个字符串 s，找到 s 中最长的回文子串。
//
// 
//
//示例 1：
//
//输入：s = "babad"
//输出："bab"
//解释："aba" 同样是符合题意的答案。
//示例 2：
//
//输入：s = "cbbd"
//输出："bb"
//示例 3：
//
//输入：s = "a"
//输出："a"
//示例 4：
//
//输入：s = "ac"
//输出："a"
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/longest-palindromic-substring
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(longestPalindrome("aaaa"))
}

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	if len(s) == 1 {
		return s
	}
	arr := [][]bool{}
	for i := 0; i < len(s); i++ {
		arr = append(arr, make([]bool, len(s)))
	}

	for i := 0; i < len(s); i++ {
		arr[i][i] = true
	}

	//for i := 0; i < len(s); i++ {
	//	for j := i ; j < len(s); j++ {
	//		fmt.Println(i,j)
	//		arr[i][j] = (i+1-j+1>=0||arr[i+1][j-1] ) && s[i] == s[j]
	//	}
	//}
	for i := 0; i < len(s); i++ {
		for j := 0; j < i; j++ {
			arr[j][i] = s[i] == s[j] && (i-j <= 2 || arr[j+1][i-1])
		}
	}

	Result := 0
	ResultStr := ""
	for i := 0; i < len(s); i++ {
		for j := 0; j < i; j++ {
			if i-j > Result && arr[i][j] {
				ResultStr = s[j : i+1]
				Result = i - j
			}
		}
	}

	for i := 0; i < len(s); i++ {
		for j := 0; j <= i; j++ {
			if i-j >= Result && arr[j][i] {
				ResultStr = s[j : i+1]
				Result = i - j
			}
		}
	}
	return ResultStr
}
