package main

import "fmt"

// 整数数组 nums 按升序排列，数组中的值 互不相同 。
//
//在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。
//
//给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。
//
// 
//
//示例 1：
//
//输入：nums = [4,5,6,7,0,1,2], target = 0
//输出：4
//示例 2：
//
//输入：nums = [4,5,6,7,0,1,2], target = 3
//输出：-1
//示例 3：
//
//输入：nums = [1], target = 0
//输出：-1
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/search-in-rotated-sorted-array
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	nus := []int{1,3,5}
	fmt.Println(search(nus,5))
}

func search(nums []int, target int) int {
	var start, end int
	end = len(nums) - 1
	for start <= end{
		half := (start + end) / 2
		if nums[half] == target{
			return half
		}
		if nums[start] == nums[half]{
			start = start + 1
		}else if nums[start] < nums[half]{
			if nums[start] <= target && target <= nums[half]{
				end = half - 1
			}else{
				start = half + 1
			}
		}else{
			if nums[half] <= target && target <= nums[end]{
				start = half + 1
			}else{
				end = half - 1
			}
		}
	}
	return - 1
}
