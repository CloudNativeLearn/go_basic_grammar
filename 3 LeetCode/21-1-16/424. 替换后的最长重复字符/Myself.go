package main

import "fmt"

// 424. 替换后的最长重复字符
// 给你一个仅由大写英文字母组成的字符串，你可以将任意位置上的字符替换成另外的字符，总共可最多替换k次。在执行上述操作后，找到包含重复字母的最长子串的长度。
//
//注意:
//字符串长度 和 k 不会超104。
//示例 1:
//
//输入:
//s = "ABAB", k = 2
//
//输出:
//4
//
//解释:
//用两个'A'替换为两个'B',反之亦然。

// 示例 2:
//
//输入:
//s = "AABABBA", k = 1
//
//输出:
//4
//
//解释:
//将中间的一个'A'替换为'B',字符串变为 "AABBBBA"。
//子串 "BBBB" 有最长重复字母, 答案为 4。

func main() {
fmt.Println(characterReplacement("AABABBA",1))
}

func characterReplacement(s string, k int) int {
	TheMap := map[int]int{}
	res := 0    //最后的结果
	maxNum := 0 //保存出现次数最多的个数

	InitMap(&TheMap)
	for i,j:=0,0;j<len(s);j++ {
		//fmt.Printf("==========i %d =======j %d====== \n",i,j)
		//fmt.Println(s[j])
		//fmt.Println(s[j]-'A')

		TheMap[int(s[j]-'A')] = TheMap[int(s[j]-'A')]+1
		maxNum = TheMax(maxNum,TheMap[int(s[j]-'A')])

		if j-i+1-maxNum>k {
			TheMap[int(s[i]-'A')] = TheMap[int(s[i]-'A')] -1
			i = i+1
			continue
		}

		res = TheMax(res,j-i+1)
	}
	return res
}

func InitMap(TheMap *map[int]int)  {
	for i:=0;i<=26;i++ {
		(*TheMap)[i] = 0
	}
}
func TheMax(one int,two int) int {
	if one>=two {
		return one
	}else {
		return two
	}
}
