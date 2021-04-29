package main

//给定两个整数数组，请交换一对数值（每个数组中取一个数值），使得两个数组所有元素的和相等。
//
//返回一个数组，第一个元素是第一个数组中要交换的元素，第二个元素是第二个数组中要交换的元素。若有多个答案，返回任意一个均可。若无满足条件的数值，返回空数组。
//
//示例:
//
//输入: array1 = [4, 1, 2, 1, 1, 2], array2 = [3, 6, 3, 3]
//输出: [1, 3]
//示例:
//
//输入: array1 = [1, 2, 3], array2 = [4, 5, 6]
//输出: []
//提示：
//
//1 <= array1.length, array2.length <= 100000
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/sum-swap-lcci
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func sum(arr []int) int {
	var ans int = 0
	for _, n := range arr {
		ans += n
	}
	return ans
}

func findSwapValues(array1 []int, array2 []int) []int {
	sum1, sum2 := sum(array1), sum(array2)
	dif := sum1 - sum2
	if dif&1 == 1 {
		return []int{}
	}
	dif >>= 1
	hash := make(map[int]struct{})
	for _, n := range array2 {
		hash[n] = struct{}{}
	}
	for _, n := range array1 {
		tar := n - dif
		if _, ok := hash[tar]; ok {
			return []int{n, tar}
		}
	}
	return []int{}
}
