package main

import "fmt"

// 给定一个数组 nums 和滑动窗口的大小 k，请找出所有滑动窗口里的最大值。
//
//示例:
//
//输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
//输出: [3,3,5,5,6,7]
//解释:
//
//  滑动窗口的位置                最大值
//---------------               -----
//[1  3  -1] -3  5  3  6  7       3
// 1 [3  -1  -3] 5  3  6  7       3
// 1  3 [-1  -3  5] 3  6  7       5
// 1  3  -1 [-3  5  3] 6  7       5
// 1  3  -1  -3 [5  3  6] 7       6
// 1  3  -1  -3  5 [3  6  7]      7

// 20 ms	6.3 MB
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	var begin = 0               // 左边坐标位置
	var ende = k - 1            // 右边坐标位置
	var canMove = len(nums) - k // 可以滑动的次数
	var ThemaxValue = -1
	var ThemaxIndex = -1

	var result []int
	for i := 0; i <= canMove; i++ {
		// 判断 最大值是否在区间中   如果在判断新加值是否比上次最大值大   如果不在重新选取区间中的最大值
		if ThemaxIndex >= begin && ThemaxIndex <= ende {
			// 新加的值为   nums[ende]
			// 判断 已知最大值 与 nums[ende] 对比
			// 1 如果比nums大 最大值依然为ThemaxValue
			// 2 如果比nums小 最大值为nums[ende]
			if ThemaxValue > nums[ende] {
				result = append(result, ThemaxValue)
			} else {
				result = append(result, nums[ende])
				ThemaxValue = nums[ende]
				ThemaxIndex = ende
			}
		} else {
			// 重新选择最大值
			ThemaxValue = nums[begin]
			ThemaxIndex = begin
			for index := begin; index <= ende; index++ {
				if nums[index] > ThemaxValue {
					ThemaxValue = nums[index]
					ThemaxIndex = index
				}
			}
			result = append(result, ThemaxValue)

		}

		begin = begin + 1
		ende = ende + 1
	}

	return result

}

func main() {
	nums := []int{-6, -10, -7, -1, -9, 9, -8, -4, 10, -5, 2, 9, 0, -7, 7, 4, -2, -10, 8, 7}
	fmt.Println(maxSlidingWindow(nums, 7))
}

