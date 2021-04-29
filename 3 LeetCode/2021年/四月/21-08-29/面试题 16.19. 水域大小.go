package main

import "fmt"

// 你有一个用于表示一片土地的整数矩阵land，该矩阵中每个点的值代表对应地点的海拔高度。若值为0则表示水域。由垂直、水平或对角连接的水域为池塘。池塘的大小是指相连接的水域的个数。编写一个方法来计算矩阵中所有池塘的大小，返回值需要从小到大排序。
//
//示例：
//
//输入：
//[
//  [0,2,1,0],
//  [0,1,0,1],
//  [1,1,0,1],
//  [0,1,0,1]
//]
//输出： [1,2,4]
//提示：
//
//0 < len(land) <= 1000
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/pond-sizes-lcci
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	//pondSizes([][]int{[[0,2,1,0],[0,1,0,1],[1,1,0,1],[0,1,0,1]]})
	fmt.Println(pondSizes([][]int{{0, 2, 1, 0}, {0, 1, 0, 1}, {1, 1, 0, 1}, {0, 1, 0, 1}}))

}



func pondSizes(land [][]int) []int {
	Lenx := len(land)
	if Lenx==0{
		return []int{}
	}
	LenY := len(land[0])
	type StarkFunc func(int, int, int) int
	var Stark StarkFunc
	Stark = func(Number int, x int, y int) int {
		if land[x][y] == 0 {
			Number = Number + 1
			land[x][y] = -1
		} else {
			return Number
		}
		//判断上下左右是否为0
		if x > 0 {
			Number = Stark(Number, x-1, y)
		}
		if y > 0 {
			Number = Stark(Number, x, y-1)
		}
		if x < Lenx-1 {
			Number = Stark(Number, x+1, y)
		}
		if y < LenY-1 {
			Number = Stark(Number, x, y+1)
		}

		// 四个顶点
		if x > 0 && y > 0 {
			Number = Stark(Number, x-1, y-1)
		}

		if x > 0 && y < LenY-1 {
			Number = Stark(Number, x-1, y+1)
		}

		if x < Lenx-1 && y < LenY-1 {
			Number = Stark(Number, x+1, y+1)
		}

		if x < Lenx-1 && y >0 {
			Number = Stark(Number, x+1, y-1)
		}
		return Number
	}
	result := []int{}
	for x := 0; x < len(land); x++ {
		for y := 0; y < len(land[x]); y++ {
			if land[x][y] == 0 {
				number := 0
				number = Stark(number, x, y)
				result = append(result, number)
			}
		}
	}

	QuickSortStark(0, len(result)-1, result)
	return result

}

// 快速排序
func QuickSortStark(l int, r int, arr []int) {
	x := (l + r) / 2
	Theleft := l
	Theright := r
	Xvalue := arr[x]
	for l < r {

		// 在右边找到比 xvalue小的值
		for arr[l] < Xvalue {
			l = l + 1
		}

		// 在左边找到比 xvalue大的值
		for arr[r] > Xvalue {
			r = r - 1
		}

		if l >= r {
			break
		}

		temp := arr[r]
		arr[r] = arr[l]
		arr[l] = temp

		if arr[l] == Xvalue {
			r = r - 1
		}
		if arr[r] == Xvalue {
			l = l + 1
		}

	}

	if l == r {
		l = l + 1
		r = r - 1
	}

	if Theleft < l {
		QuickSortStark(Theleft, r, arr)
	}
	if Theright > r {
		QuickSortStark(l, Theright, arr)
	}
}
