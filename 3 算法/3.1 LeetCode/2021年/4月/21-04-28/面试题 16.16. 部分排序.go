package main

import "sort"

// 给定一个整数数组，编写一个函数，找出索引m和n，只要将索引区间[m,n]的元素排好序，整个数组就是有序的。注意：n-m尽量最小，也就是说，找出符合条件的最短序列。函数返回值为[m,n]，若不存在这样的m和n（例如整个数组是有序的），请返回[-1,-1]。
//
//示例：
//
//输入： [1,2,4,7,10,11,7,12,6,7,16,18,19]
//输出： [3,9]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/sub-sort-lcci
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	subSort([]int{1,2,4,7,10,11,7,12,6,7,16,18,19})
}
func subSort(array []int) []int {
	if len(array) <=  1{
		return []int{-1 , -1}
	}
	var res []int
	for i := 0 ; i < len(array) ; i++{
		res = append(res , array[i])
	}
	sort.Ints(res)
	n := len(array)
	l , r := -1 ,-1
	for i := 0 ; i < n ; i++{
		if res[i] != array[i]{
			l = i
			break
		}
	}
	for i := n - 1 ; i > 0 ; i--{
		if res[i] != array[i]{
			r = i
			break
		}
	}
	return []int{l , r}
}
