package main

import "fmt"

// 3 x 3 的幻方是一个填充有从 1 到 9 的不同数字的 3 x 3 矩阵，其中每行，每列以及两条对角线上的各数之和都相等。
//
//给定一个由整数组成的 grid，其中有多少个 3 × 3 的 “幻方” 子矩阵？（每个子矩阵都是连续的）。
//
// 
//
//示例：
//
//输入: [[4,3,8,4],
//      [9,5,1,9],
//      [2,7,6,2]]
//输出: 1
//解释:
//下面的子矩阵是一个 3 x 3 的幻方：
//438
//951
//276
//
//而这一个不是：
//384
//519
//762
//
//总的来说，在本示例所给定的矩阵中只有一个 3 x 3 的幻方子矩阵。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/magic-squares-in-grid
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 3*3 的表格
func main() {
	fmt.Println(numMagicSquaresInside([][]int{{4, 3, 8, 4}, {9, 5, 1, 9}, {2, 7, 6, 2}}))
}
func numMagicSquaresInside(grid [][]int) int {
	// 类似前缀和
	heng := [][]int{{1}}
	shu := [][]int{{1}}

	//OneDui := [][]int{{1}}
	//TwoDui := [][]int{{1}}

	for i:=0;i<len(grid);i++{
		for j:=0;j<len(grid[i]);j++{
			if i==0 || j==0{
				heng[i][j] = grid[i][j]
			}else {
				heng[i][j] = grid[i][j]*grid[i][j-1]
			}
		}
	}


	for i:=0;i<len(grid[0]);i++{
		for j:=0;j<len(grid);j++{
			if i==0 || j==0{
				shu[i][j] = grid[i][j]
			}else {
				shu[i][j] = grid[i][j]*grid[j-1][i]
			}
		}
	}

	// f
	result := 0
	for i:=0;i+3<len(grid);i++{
		for j:=0;j+3<len(grid[i]);j++{
			whetherHengshu := (shu[i][j]/shu[i+3][j] == shu[i+1][j]/shu[i+4][j]) && (shu[i+1][j]/shu[i+4][j] == shu[i+2][j]/shu[i+5][j]) && (shu[i+2][j]/shu[i+5][j] == shu[i+3][j]/shu[i+6][j])
			whetherShu := (shu[i][j]/shu[i][j+3] == shu[i][j+1]/shu[i][j+4]) && (shu[i][j+1]/shu[i][j+4] == shu[i][j+2]/shu[i][j+5]) && (shu[i][j+2]/shu[i][j+5] == shu[i][j+3]/shu[i][j+6])

			one := shu[i][j] * shu[i+1][j+1] * shu[i+2][j+2]
			two := shu[i][j+2] * shu[i+1][j+1] *  shu[i+2][j]

			three := one==two && whetherHengshu && whetherShu && shu[i][j]/shu[i+3][j]==one
			if three{
				result = result+1
			}
		}
	}
	return  result

}

// 判断数组里面的值是否相等
func Whether(arr []int) bool {
	x := arr[0]
	for _, v := range arr {
		if x != v {
			return false
		}
	}
	return true
}
