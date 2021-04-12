package main

// 给定一个包含非负整数的数组，你的任务是统计其中可以组成三角形三条边的三元组个数。
//
//示例 1:
//
//输入: [2,2,3,4]
//输出: 3
//解释:
//有效的组合是:
//2,3,4 (使用第一个 2)
//2,3,4 (使用第二个 2)
//2,2,3
//注意:
//
//数组长度不超过1000。
//数组里整数的范围为 [0, 1000]。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/valid-triangle-number
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {

}

func triangleNumber(nums []int) int {
	//	将数组排序
	triangleNumber(nums)
	l := 0
	r := 0
	res := 0
	numbers := 0
	for r <len(nums) && l<len(nums){
		if  {

		}
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
