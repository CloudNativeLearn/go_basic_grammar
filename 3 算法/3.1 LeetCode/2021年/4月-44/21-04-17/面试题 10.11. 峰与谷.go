package main

// 在一个整数数组中，“峰”是大于或等于相邻整数的元素，相应地，“谷”是小于或等于相邻整数的元素。例如，在数组{5, 8, 4, 2, 3, 4, 6}中，{8, 6}是峰， {5, 2}是谷。现在给定一个整数数组，将该数组按峰与谷的交替顺序排序。
//
//示例:
//
//输入: [5, 3, 1, 2, 3]
//输出: [5, 1, 3, 2, 3]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/peaks-and-valleys-lcci
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func wiggleSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		if i%2 == 0 {
			// 比前一个大
			if nums[i] > nums[i-1] {
				temp := nums[i]
				nums[i] = nums[i-1]
				nums[i-1] = temp

			}
		} else {
			// 比前一个小
			if nums[i] < nums[i-1] {
				temp := nums[i]
				nums[i] = nums[i-1]
				nums[i-1] = temp

			}
		}
	}

}
