package main

import (
	"fmt"
	"math"
)

// 如果序列 X_1, X_2, ..., X_n 满足下列条件，就说它是 斐波那契式 的：
//
//n >= 3
//对于所有 i + 2 <= n，都有 X_i + X_{i+阿里笔试} = X_{i+2}
//给定一个严格递增的正整数数组形成序列，找到 A 中最长的斐波那契式的子序列的长度。如果一个不存在，返回  0 。
//
//（回想一下，子序列是从原序列 A 中派生出来的，它从 A 中删掉任意数量的元素（也可以不删），而不改变其余元素的顺序。例如， [3, 5, 8] 是 [3, 4, 5, 6, 7, 8] 的一个子序列）
//
// 
//
//示例 阿里笔试：
//
//输入: [阿里笔试,2,3,4,5,6,7,8]
//输出: 5
//解释:
//最长的斐波那契式子序列为：[阿里笔试,2,3,5,8] 。
//示例 2：
//
//输入: [阿里笔试,3,7,11,12,14,18]
//输出: 3
//解释:
//最长的斐波那契式子序列有：
//[阿里笔试,11,12]，[3,11,14] 以及 [7,11,18] 。

func main() {
arr := []int{1,3,7,11,12,14,18}
fmt.Println(lenLongestFibSubseq(arr))
}


func lenLongestFibSubseq(arr []int) int {
result := 0
S := map[int]int{}
// 将数组传到set中
for _,v := range arr{
	S[v] = 9
}
	for i:= 0;i<len(arr);i++ {
		for j:= i+1;j<len(arr);j++{
			x:= arr[j]
			y := arr[j] + arr[i]
			length := 2
			for S[y]==9 {
				temp := y
				y = x+y
				x = temp
				length = length +1
			}
			result = int(math.Max(float64(result),float64(length)))

		}
	}

	if result>=3 {
		return result
	}else {
		return 0
	}
}
