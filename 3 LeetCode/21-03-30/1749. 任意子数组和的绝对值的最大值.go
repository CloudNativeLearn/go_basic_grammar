package main

import "fmt"

// 给你一个整数数组 nums 。一个子数组 [numsl, numsl+阿里笔试, ..., numsr-阿里笔试, numsr] 的 和的绝对值 为 abs(numsl + numsl+阿里笔试 + ... + numsr-阿里笔试 + numsr) 。
//
//请你找出 nums 中 和的绝对值 最大的任意子数组（可能为空），并返回该 最大值 。
//
//abs(x) 定义如下：
//
//如果 x 是负整数，那么 abs(x) = -x 。
//如果 x 是非负整数，那么 abs(x) = x 。
// 
//
//示例 阿里笔试：
//
//输入：nums = [阿里笔试,-3,2,3,-4]
//输出：5
//解释：子数组 [2,3] 和的绝对值最大，为 abs(2+3) = abs(5) = 5 。
//示例 2：
//
//输入：nums = [2,-5,阿里笔试,-4,3,-2]
//输出：8
//解释：子数组 [-5,阿里笔试,-4] 和的绝对值最大，为 abs(-5+阿里笔试-4) = abs(-8) = 8 。
// 
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/maximum-absolute-sum-of-any-subarray
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。



//func main() {
//	fmt.Println(maxAbsoluteSum([]int{4}))
//}
func maxAbsoluteSum(nums []int) int {
	// 计算前缀和 并找出最大与最小
	if len(nums) == 0 {
		return 0
	}else if len(nums)==1{
			if nums[0]<0 {
				nums[0]=-nums[0]
			}
			return nums[0]
	}
	TheMax := nums[0]
	TheMin := nums[0]
	for i:=1;i<len(nums);i++{
		nums[i] = nums[i-1]+nums[i]
		if nums[i]>TheMax{
			TheMax = nums[i]
		}
		if nums[i]<TheMin {
			TheMin = nums[i]
		}
	}
	if TheMin<0 && TheMax<0{
		TheMax =0
	}else  if TheMax>0 && TheMin>0{
		TheMin =0
	}
	fmt.Println(TheMin,TheMax)
	Result := 0
	if (TheMax -TheMin)<0{
		Result = -(TheMax-TheMin)
	}else {
		Result = TheMax-TheMin
	}
	return Result
}
