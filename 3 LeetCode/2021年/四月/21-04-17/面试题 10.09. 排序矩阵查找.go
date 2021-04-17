package main

// 给定M×N矩阵，每一行、每一列都按升序排列，请编写代码找出某元素。
//
//示例:
//
//现有矩阵 matrix 如下：
//
//[
//  [1,   4,  7, 11, 15],
//  [2,   5,  8, 12, 19],
//  [3,   6,  9, 16, 22],
//  [10, 13, 14, 17, 24],
//  [18, 21, 23, 26, 30]
//]
//给定 target = 5，返回 true。
//
//给定 target = 20，返回 false。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/sorted-matrix-search-lcci
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func searchMatrix(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		if len(matrix[i])!=0&&  matrix[i][0]<=target&&matrix[i][len(matrix[i])-1]>=target{
			for j:=0;j<len(matrix[i]);j++{
				if matrix[i][j] == target{
					return true
				}
			}
		}
	}
	return false
}
