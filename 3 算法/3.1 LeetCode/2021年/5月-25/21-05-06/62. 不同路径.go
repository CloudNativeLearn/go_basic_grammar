package main

import "fmt"

// 一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
//
//机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
//
//问总共有多少条不同的路径？
//
// 
//
//示例 1：
//
//
//输入：m = 3, n = 7
//输出：28
//示例 2：
//
//输入：m = 3, n = 2
//输出：3
//解释：
//从左上角开始，总共有 3 条路径可以到达右下角。
//1. 向右 -> 向下 -> 向下
//2. 向下 -> 向下 -> 向右
//3. 向下 -> 向右 -> 向下
//示例 3：
//
//输入：m = 7, n = 3
//输出：28
//示例 4：
//
//输入：m = 3, n = 3
//输出：6
// 
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/unique-paths
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(uniquePaths(3,7))
}


func uniquePaths(m int, n int) int {
	if m==0||n==0{
		return 0
	}
	arr := make([][]int, m)
	for i := 0; i < m; i++ {
		//arr = append(arr, make([]int, n))
		arr[i] = make([]int, n)
	}

	arr[0][0] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			result := 0
			if i>=1{
				result = result + arr[i-1][j]
			}
			if j>=1{
				result = result + arr[i][j-1]
			}
			if !(j==0 && i==0){
				arr[i][j] = result
			}
		}
	}


	return arr[m-1][n-1]

}
