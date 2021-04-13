package main

import "fmt"
// // 计算一堆任务的时间和（多个并发任务占用的时间只算一次，具体看下面例子）
////输入一个任务数组，每个任务有开始和结束时间 ， 开始> 结束时间 （时间为 long类型）
////例如
////a.输入 一个任务 [1,4], 要求求的时间就是 4-1=3
////b.输入 两个任务 [1,4],[6,8]   要求求的时间就是5
////c.输入 两个任务 [1,4],[2,5]   要求求的时间就是4 .(并发的时间不会算两遍)
////最后一个例子
////d. 输入3个任务 [1,4],[6,8] ,[2,5] 返回 6.
////java 相关方法，也可用任何其他语言
////Task 可以调用的方法 getStart(), getEnd()  分别返回开始和结束时间


func main() {
	arr := [][]int{{1,9},{3,4},{4,9}}
	fmt.Println(getSum(arr))
}

func getSum(arr [][]int) int {
	if len(arr) == 0 {
		return 0
	}
	if len(arr) == 1 {
		return arr[0][1] - arr[0][0]
	}
	QuickSortStark(0, len(arr)-1, arr)
	fmt.Println(arr)
	result := 0
	begin := arr[0][0]
	end := arr[0][1]
	for i := 1; i < len(arr); i++ {
		if arr[i][0] > end {
			result = result + end - begin
			begin = arr[i][0]
			end = arr[i][1]
		} else {
			if end >= arr[i][1] {
				end = end
			} else {
				end = arr[i][1]
			}
		}
	}

	return result + end - begin
}

// 快速排序
func QuickSortStark(l int, r int, arr [][]int) {
	x := (l + r) / 2
	Theleft := l
	Theright := r
	Xvalue := arr[x][0]
	for l < r {

		// 在右边找到比 xvalue小的值
		for arr[l][0] < Xvalue {
			l = l + 1
		}

		// 在左边找到比 xvalue大的值
		for arr[r][0] > Xvalue {
			r = r - 1
		}

		if l >= r {
			break
		}

		temp := arr[r]
		arr[r] = arr[l]
		arr[l] = temp

		if arr[l][0] == Xvalue {
			r = r - 1
		}
		if arr[r][0] == Xvalue {
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
