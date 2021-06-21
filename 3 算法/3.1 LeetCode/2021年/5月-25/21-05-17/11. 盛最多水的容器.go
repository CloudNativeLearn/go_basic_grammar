package main

// 给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0) 。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
//
//说明：你不能倾斜容器。
//
// 
//
//示例 1：
//
//
//
//输入：[1,8,6,2,5,4,8,3,7]
//输出：49
//解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
//示例 2：
//
//输入：height = [1,1]
//输出：1
//示例 3：
//
//输入：height = [4,3,2,1,4]
//输出：16
//示例 4：
//
//输入：height = [1,2,1]
//输出：2
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/container-with-most-water
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	result := 0
	for l != r {
		result = GettMax(result, (r-l)*GetMin(height[l], height[r]))
		if height[l]>height[r]{
			r = r-1
		}else {
			l = l+1
		}
	}
	return result
}

func GettMax(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func GetMin(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
