package main

import "fmt"

//
// 给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的 连续 子数组，并返回其长度。如果不存在符合条件的子数组，返回 0。
// 示例：
//输入：s = 7, nums = [2,3,1,2,4,3]
//输出：2
//解释：子数组 [4,3] 是该条件下的长度最小的子数组。



func main() {
	arry := []int{2, 3, 1, 2, 4, 3}
	s := 7
	fmt.Println(minSubArrayLen(s,arry))

}

func minSubArrayLen(s int, nums []int) int {
	l := 0
	r := 0
	arry := &nums
	result := 0
	maxLen := len(*arry)
	if maxLen == 0{
		return 0
	}
	for {
		if r == l && r > 0 {
			break
		}
		y := (*arry)[l:r]
		sum := GetTheNUmbera(&y)

		if sum >= s {
			// 结果比需要的值大
			result = TheMinNUmber(result, r-l)

			l = l + 1
		} else {
			if r == maxLen {
				l = l + 1
			} else {
				r = r + 1
			}

		}

	}

	return result


}

func GetTheNUmbera(arr *[]int) int {
	sum := 0
	for _, v := range *arr {
		sum = sum + v
	}
	return sum
}

// 判断
func TheMinNUmber(one int, two int) int {
	if two <= 0 {
		return one
	} else if one <= 0 {
		return two
	} else if one < two {
		return one
	} else {
		return two
	}
}
