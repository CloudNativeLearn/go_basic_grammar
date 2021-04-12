package main

// 编写一种算法，若M × N矩阵中某个元素为0，则将其所在的行与列清零。
//
// 
//
//示例 1：
//
//输入：
//[
//  [1,1,1],
//  [1,0,1],
//  [1,1,1]
//]
//输出：
//[
//  [1,0,1],
//  [0,0,0],
//  [1,0,1]
//]
//示例 2：
//
//输入：
//[
//  [0,1,2,0],
//  [3,4,5,2],
//  [1,3,1,5]
//]
//输出：
//[
//  [0,0,0,0],
//  [0,4,5,0],
//  [0,3,1,0]
//]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/zero-matrix-lcci
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。



func setZeroes(matrix [][]int)  {
	row := map[int]int{} // 行
	col := map[int]int{} // 列
	for index1,value1 := range matrix{
		for index2,value2 := range value1{
			if value2==0{
				row[index1] =0
				col[index2] = 0
			}
		}
	}

	for k,_ := range row{
		for j:=0;j<len(matrix[k]);j++{
			matrix[k][j] =0
		}
	}

	for k,_ := range col{
		for j:=0;j<len(matrix);j++{
			matrix[j][k] =0
		}
	}
}
