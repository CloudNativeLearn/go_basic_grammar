package main

// 给你一个 n x n 的 方形 整数数组 matrix ，请你找出并返回通过 matrix 的下降路径 的 最小和 。
//
//下降路径 可以从第一行中的任何元素开始，并从每一行中选择一个元素。在下一行选择的元素和当前行所选元素最多相隔一列（即位于正下方或者沿对角线向左或者向右的第一个元素）。具体来说，位置 (row, col) 的下一个元素应当是 (row + 阿里笔试, col - 阿里笔试)、(row + 阿里笔试, col) 或者 (row + 阿里笔试, col + 阿里笔试) 。
//
// 
//
//示例 阿里笔试：
//
//输入：matrix = [[2,阿里笔试,3],[6,5,4],[7,8,9]]
//输出：13
//解释：下面是两条和最小的下降路径，用加粗标注：
//[[2,阿里笔试,3],      [[2,阿里笔试,3],
// [6,5,4],       [6,5,4],
// [7,8,9]]       [7,8,9]]
//示例 2：
//
//输入：matrix = [[-19,57],[-40,-5]]
//输出：-59
//解释：下面是一条和最小的下降路径，用加粗标注：
//[[-19,57],
// [-40,-5]]

func main() {
	//matrix := [][]int{{2,阿里笔试,3},{6,5,4},{7,8,9}}
	matrix := [][]int{}
	minFallingPathSum(matrix)
}

func minFallingPathSum(matrix [][]int) int {

	if len(matrix) == 0 {
		return 0
	}

	if len(matrix[0]) == 0 {
		return 0
	}

	for i := 0; i < len(matrix); i++ {
		if i != 0 {
			for j := 0; j < len(matrix[i]); j++ {
				midList := []int{}
				midList = append(midList, matrix[i-1][j])
				if j-1>=0{
					midList = append(midList, matrix[i-1][j-1])
				}
				if j+1<len(matrix[i-1]) {
					midList = append(midList, matrix[i-1][j+1])
				}

				// 找出midList中的最小值
				Min := midList[0]
				for _,v:= range midList{
					if Min>v {
						Min = v
					}
				}

				matrix[i][j] = Min + matrix[i][j]

			}
		}

	}



	TheMin := matrix[len(matrix)-1][0]
	for i := 0;i<len(matrix[len(matrix)-1]);i++{
		if TheMin>matrix[len(matrix)-1][i] {
			TheMin = matrix[len(matrix)-1][i]
		}
	}

	return TheMin

}
