package main

//给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
//
//说明：每次只能向下或者向右移动一步。
//
// 
//
//示例 1：
//
//
//输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
//输出：7
//解释：因为路径 1→3→1→1→1 的总和最小。
//示例 2：
//
//输入：grid = [[1,2,3],[4,5,6]]
//输出：12
// 
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/minimum-path-sum
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func minPathSum(grid [][]int) int {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			a := 9223372036854775807
			if i >= 1 {
				// 上
				a  = GetMin(a,grid[i-1][j])
			}

			if j >= 1 {
				// 左
				a = GetMin(a,grid[i][j-1])
			}

			if !(i==0 && j==0){
				grid[i][j] = grid[i][j] + a
			}

		}
	}

	return grid[len(grid)-1][len(grid[len(grid)-1])-1]

}

func GetMin(a int,b int) int {
	if a>b{
		return b
	}else {
		return a
	}
}
