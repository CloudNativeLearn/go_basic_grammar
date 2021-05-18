package main

// 给定两个单词 word1 和 word2，找到使得 word1 和 word2 相同所需的最小步数，每步可以删除任意一个字符串中的一个字符。
//
// 
//
//示例：
//
//输入: "sea", "eat"
//输出: 2
//解释: 第一步将"sea"变为"ea"，第二步将"eat"变为"ea"
// 
//
//提示：
//
//给定单词的长度不超过500。
//给定单词中的字符只含有小写字母。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/delete-operation-for-two-strings
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。




func minDistance(word1 string, word2 string) int {
	result := [][]int{}
	for i := 0; i <= len(word1); i++ {
		result = append(result, make([]int, len(word2)+1))
		for j := 0; j <= len(word2); j++ {
			if i == 0 {
				if j > 0 {
					result[0][j] = result[0][j-1] + 1
				}
			}else {
				if j==0{
					result[i][0] = result[i-1][0] +1
				}else {
					// 如果两个值相等

					if word1[i-1] == word2[j-1]{
						result[i][j] = result[i-1][j-1]
					}else {
						result[i][j] = GetMin(result[i-1][j],result[i][j-1]) +1
					}
					// 如果两个值不相等
				}
			}
		}

	}


	return result[len(word1)][len(word2)]
}

func GetMin(a int,b int) int {
	if a<b {
		return a
	}
	return b
}
//    “”  e  a   t
// "" 0   1  2   3
// s  1   2  3   4
// e  2   1  4   5
// a  3   4  1   6
