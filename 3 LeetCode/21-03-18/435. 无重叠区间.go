package main

import "fmt"

// 给定一个区间的集合，找到需要移除区间的最小数量，使剩余区间互不重叠。
//
//注意:
//
//可以认为区间的终点总是大于它的起点。
//区间 [1,2] 和 [2,3] 的边界相互“接触”，但没有相互重叠。
//示例 1:
//
//输入: [ [1,2], [2,3], [3,4], [1,3] ]
//
//输出: 1
//
//解释: 移除 [1,3] 后，剩下的区间没有重叠。
//示例 2:
//
//输入: [ [1,2], [1,2], [1,2] ]
//
//输出: 2
//
//解释: 你需要移除两个 [1,2] 来使剩下的区间没有重叠。
//示例 3:
//
//输入: [ [1,2], [2,3] ]
//
//输出: 0
//
//解释: 你不需要移除任何区间，因为它们已经是无重叠的了。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/non-overlapping-intervals
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
arr := [][]int{{1,2},{2,3},{3,4},{1,3}}
fmt.Println(eraseOverlapIntervals(arr))
fmt.Println(arr)
}
func eraseOverlapIntervals(intervals [][]int) int {

	// 当只有一个数据的时候  返回0
	result :=1
	if len(intervals)==1 || len(intervals)==0 {
		return 0
	}else {
		QuickSortStark(0,len(intervals)-1,intervals)
		fmt.Println(intervals)
		begin :=0
		for i:=1;i<len(intervals);i++ {
			if intervals[begin][len(intervals[begin])-1] <= intervals[i][0] {
				begin = i
				result = result+1
			}
		}
	}
	fmt.Println(result)


	return len(intervals)-result


}



func QuickSortStark(l int,r int,arr [][]int)  {
	x := (l+r)/2
	Theleft := l
	Theright := r
	Xvalue := arr[x][len(arr[x])-1]
	for l<r{

		// 在右边找到比 xvalue小的值
		for arr[l][len(arr[l])-1]<Xvalue{
			l = l+1
		}

		// 在左边找到比 xvalue大的值
		for  arr[r][len(arr[r])-1]>Xvalue {
			r = r-1
		}

		if l>=r {
			break
		}

		temp := arr[r]
		arr[r] = arr[l]
		arr[l] = temp

		if arr[l][len(arr[l])-1] == Xvalue{
			r = r-1
		}
		if arr[r][len(arr[r])-1] == Xvalue{
			l = l+1
		}

	}

	if l ==r {
		l = l+1
		r = r-1
	}

	if Theleft<l {
		QuickSortStark(Theleft,r,arr)
	}
	if Theright>r {
		QuickSortStark(l,Theright,arr)
	}
}