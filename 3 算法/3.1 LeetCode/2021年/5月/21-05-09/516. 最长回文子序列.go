package main

// 给定一个字符串 s ，找到其中最长的回文子序列，并返回该序列的长度。可以假设 s 的最大长度为 1000 。
//
// 
//
//示例 1:
//输入:
//
//"bbbab"
//输出:
//
//4
//一个可能的最长回文子序列为 "bbbb"。
//
//示例 2:
//输入:
//
//"cbbd"
//输出:
//
//2
//一个可能的最长回文子序列为 "bb"。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/longest-palindromic-subsequence
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	longestPalindromeSubseq("bbbab")
}

func longestPalindromeSubseq(s string) int {
	n := len(s)
	if len(s)==0{
		return 0
	}
	if len(s)==1{
		return 1
	}
	arr := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		arr[i] = make([]int, len(s))
	}

	result := 0
	for i := n-1; i >=0; i-- {
		arr[i][i]=1
		for j := i+1; j < len(s); j++ {

			if s[i]==s[j]{
				arr[i][j] = arr[i+1][j-1]+2
			}else {
				arr[i][j] = GetTheMax(arr[i+1][j],arr[i][j-1])
			}
			result =  GetTheMax(result,arr[i][j])
		}
	}

	//fmt.Println(arr)
	return result
}

func GetTheMax(a int, b int)int  {
	if a>b{
		return a
	}else {
		return b
	}
}