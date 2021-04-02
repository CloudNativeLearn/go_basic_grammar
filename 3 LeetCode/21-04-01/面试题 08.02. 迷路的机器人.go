package main

import "fmt"

// 设想有个机器人坐在一个网格的左上角，网格 r 行 c 列。机器人只能向下或向右移动，但不能走到一些被禁止的网格（有障碍物）。设计一种算法，寻找机器人从左上角移动到右下角的路径。
//
//
//
//网格中的障碍物和空位置分别用 1 和 0 来表示。
//
//返回一条可行的路径，路径由经过的网格的行号和列号组成。左上角为 0 行 0 列。如果没有可行的路径，返回空数组。
//
//示例 1:
//
//输入:
//[
//  [0,0,0],
//  [0,1,0],
//  [0,0,0]
//]
//输出: [[0,0],[0,1],[0,2],[1,2],[2,2]]
//解释:
//输入中标粗的位置即为输出表示的路径，即
//0行0列（左上角） -> 0行1列 -> 0行2列 -> 1行2列 -> 2行2列（右下角）
//说明：r 和 c 的值均不超过 100。
//
//通过次数8,191提交次数23,018
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/robot-in-a-grid-lcci
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
x := pathWithObstacles([][]int{{1}})
fmt.Println(x)
}
//
//func pathWithObstacles(obstacleGrid [][]int) [][]int {
//	if len(obstacleGrid)==0 {
//		return [][]int{}
//	}
//	if len(obstacleGrid)==1 {
//		if len(obstacleGrid[0])==1{
//			if obstacleGrid[0][0]==0 {
//				return [][]int{{0,0}}
//			}else {
//				return [][]int{}
//			}
//		}
//
//	}
//	LenX := len(obstacleGrid)
//	LenY := len(obstacleGrid[0])
//	Result := [][]int{}
//	y,_ := PathStark(&obstacleGrid,0,0,LenX,LenY,&Result)
//	return *y
//}
//
//func PathStark(arr *[][]int, x int, y int, LenX int, LenY int, Result *[][]int) (*[][]int, bool) {
//	xWhether := false
//	if (*arr)[x][y]==1 {
//		return Result,false
//	}
//	*Result = append(*Result, []int{x,y})
//
//	if x+1 < LenX {
//		if (*arr)[x+1][y] != 1 {
//			NewArr := make([][]int, len(*Result))
//			copy(NewArr, *Result)
//			//NewArr = append(NewArr, []int{x, y})
//			xarr := &[][]int{}
//			xarr, xWhether = PathStark(arr, x+1, y, LenX, LenY, &NewArr)
//			if xWhether {
//				Result = xarr
//				return xarr,true
//			}
//		}
//	}
//
//	yWhether := false
//	if y+1 < LenY {
//		if (*arr)[x][y+1] != 1 {
//			NewArr := make([][]int, len(*Result))
//			copy(NewArr, *Result)
//			//NewArr = append(NewArr, []int{x, y})
//			xarr := &[][]int{}
//			xarr,yWhether = PathStark(arr, x, y+1, LenX, LenY, &NewArr)
//			if yWhether {
//				Result = xarr
//				return xarr,true
//			}
//		}
//	}
//
//	if x == LenX-1 && y == LenY-1 {
//
//		return Result, true
//	}
//
//	return &[][]int{},false
//}



func pathWithObstacles(obstacleGrid [][]int) [][]int {
	var (
		result [][]int
		rowNum = len(obstacleGrid) - 1
		colNum = len(obstacleGrid[0]) - 1
		dfs    func([][]int)
	)
	dfs = func(path [][]int) {
		if len(result) != 0 {
			return
		}
		curNode := path[len(path)-1]
		row, col := curNode[0], curNode[1]
		if obstacleGrid[row][col] != 0 {
			return
		}

		if row == rowNum && col == colNum {
			result = make([][]int, rowNum+colNum+1)
			copy(result, path)
		}

		obstacleGrid[row][col] = 1
		// 向下走
		if row < rowNum {
			dfs(append(path, []int{row + 1, col}))
		}
		// 向右走
		if col < colNum {
			dfs(append(path, []int{row, col + 1}))
		}
	}
	dfs([][]int{{0, 0}})
	return result
}
