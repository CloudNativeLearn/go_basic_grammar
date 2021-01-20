package main

// 给定两个字符串s1, s2，找到使两个字符串相等所需删除字符的ASCII值的最小和。
//
//示例 1:
//
//输入: s1 = "sea", s2 = "eat"
//输出: 231
//解释: 在 "sea" 中删除 "s" 并将 "s" 的值(115)加入总和。
//在 "eat" 中删除 "t" 并将 116 加入总和。
//结束时，两个字符串相等，115 + 116 = 231 就是符合条件的最小和。
//示例2:
//
//输入: s1 = "delete", s2 = "leet"
//输出: 403
//解释: 在 "delete" 中删除 "dee" 字符串变成 "let"，
//将 100[d]+101[e]+101[e] 加入总和。在 "leet" 中删除 "e" 将 101[e] 加入总和。
//结束时，两个字符串都等于 "let"，结果即为 100+101+101+101 = 403 。
//如果改为将两个字符串转换为 "lee" 或 "eet"，我们会得到 433 或 417 的结果，比答案更大。
//注意:
//
//0 < s1.length, s2.length <= 1000。
//所有字符串中的字符ASCII值在[97, 122]之间。

func main() {

}

func minimumDeleteSum(s1 string, s2 string) int {
	s1Len := len(s1)
	s2Len := len(s2)
	db:= make([][]int,s1Len+1)
	for i:=0;i<=s1Len;i++ {
		db[i] = make([]int,s2Len+1)
	}
	db[0][0] = 0
	for i:=1;i<=s1Len;i++ {
		db[i][0] = db[i-1][0] + int(s1[i-1])
	}

	for i:=1;i<=s2Len;i++ {
		db[0][i] = db[0][i-1] + int(s2[i-1])
	}

	for i:=1;i<=s1Len;i++ {
		for j:=1;j<=s2Len;j++ {
			if s1[i-1] == s2[j-1] {
				db[i][j] = db[i-1][j-1]
			}else {
				temp1 := db[i-1][j] + int(s1[i-1])
				temp2 := db[i][j-1] + int(s2[j-1])
				if temp1 < temp2{
					db[i][j] = temp1
				}else {
					db[i][j] = temp2
				}
			}




		}
	}
return db[s1Len][s2Len]

}

