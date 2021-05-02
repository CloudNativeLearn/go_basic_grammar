package main

import "fmt"

// 给定两个整数数组a和b，计算具有最小差绝对值的一对数值（每个数组中取一个值），并返回该对数值的差
//
// 
//
//示例：
//
//输入：{1, 3, 15, 11, 2}, {23, 127, 235, 19, 8}
//输出：3，即数值对(11, 8)
// 
//
//提示：
//
//1 <= a.length, b.length <= 100000
//-2147483648 <= a[i], b[i] <= 2147483647
//正确结果在区间 [0, 2147483647] 内
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/smallest-difference-lcci
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(smallestDifference([]int{1, 3, 15, 11, 2},[]int{23, 127, 235, 19, 8}))
}

func smallestDifference(a []int, b []int) int {
	QuickSortStark(0, len(a)-1, a)
	QuickSortStark(0, len(b)-1, b)
	l := 0
	r := 0
	result := 2147483647
	for l < len(a) && r < len(b) {

		for  {

			if r+1==len(b)-1{
				break
			}
			if r==len(b)-1{
				break
			}
			if a[l]<=b[r+1]{
				break
			}
			r = r+1
		}
		if r+1<len(b){
			result2 := GetMin(ChaAbs(a[l],b[r]),ChaAbs(a[l],b[r+1]))
			result  = GetMin(result2,result)
		}else {
			result2 := ChaAbs(a[l],b[r])
			result = GetMin(result2,result)
		}
		//result = GetMin(ChaAbs(a[l],b[r]),ChaAbs(a[l],b[r+1]))
		l = l+1
	}
	return  result
}

// 获得更小的数
func GetMin(a int,b int)int  {
	if a<b{
		return a
	}else {
		return b
	}
}

// 两个数差的绝对值
func ChaAbs(a int, b int) int {
	c := a - b
	if c < 0 {
		return -c
	} else {
		return c
	}
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