package main

import (
	"fmt"
	"strconv"
)

// 给定一个放有字符和数字的数组，找到最长的子数组，且包含的字符和数字的个数相同。
//
//返回该子数组，若存在多个最长子数组，返回左端点下标值最小的子数组。若不存在这样的数组，返回一个空数组。
//
//示例 1:
//
//输入: ["A","1","B","C","D","2","3","4","E","5","F","G","6","7","H","I","J","K","L","M"]
//
//输出: ["A","1","B","C","D","2","3","4","E","5","F","G","6","7"]
//示例 2:
//
//输入: ["A","A"]
//
//输出: []
//提示：
//
//array.length <= 100000
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/find-longest-subarray-lcci
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(findLongestSubarray([]string{"A", "1", "B", "C", "D", "2", "3", "4", "E", "5", "F", "G", "6", "7", "H", "I", "J", "K", "L", "M"}))

}

func findLongestSubarray(array []string) []string {
	ThNumberArr := []int{}
	for _, v := range array {
		_, error := strconv.Atoi(v)
		if error == nil {
			if len(ThNumberArr) != 0 {
				ThNumberArr = append(ThNumberArr, ThNumberArr[len(ThNumberArr)-1]+1)
			} else {
				ThNumberArr = append(ThNumberArr, 1)
			}
		} else {
			if len(ThNumberArr) != 0 {
				ThNumberArr = append(ThNumberArr, ThNumberArr[len(ThNumberArr)-1])
			} else {
				ThNumberArr = append(ThNumberArr, 0)
			}
		}
	}

	//result := [][]int{}
	for i := 0; i < len(array); i++ {
		for j := len(array) - 1; j >= 0; j-- {

			index := 0
			for m := j; m < len(array); m++ {

				lNumber := ThNumberArr[i+index]
				lLetter := i + 1 - lNumber
				rNumber := ThNumberArr[m]
				rLetter := j + 1 - rNumber

				// 判断 i 这个数字的状态
				_, error := strconv.Atoi(array[i+index])
				if error == nil {
					// 数字
					if lNumber-1-rNumber == lLetter-rLetter {
						//result = append(result, []int{i, j + 1})
						//break
						return array[i+index : m+1]
					}
				} else {
					if lNumber-rNumber == lLetter-1-rLetter {
						//result = append(result, []int{i, j + 1})
						//break
						return array[i+index : m+1]
					}
				}

				index = index + 1

			}

		}
	}

	return []string{}
}

func GetMax(arr [][]int) []int {
	x := arr[0]
	for _, v := range arr {
		if v[1]-v[0] > x[1]-x[0] {
			x = v
		}
	}
	return x
}

//func findLongestSubarray(array []string) []string {
//	TheLetter := 0
//	TheNumber := 0
//
//	for _, v := range array {
//		_, error := strconv.Atoi(v)
//		if error == nil {
//			TheNumber = TheNumber + 1
//		} else {
//			TheLetter = TheLetter + 1
//		}
//	}
//
//
//
//
//	if TheLetter == TheNumber {
//		return array
//	}
//
//
//
//
//
//	l := 0
//	r := len(array) - 1
//
//
//	for l != r {
//		_, error := strconv.Atoi(array[r])
//		// 判断 r-1
//		if error == nil {
//			TheNumber = TheNumber - 1
//		} else {
//			TheLetter = TheLetter - 1
//		}
//
//		if TheLetter == TheNumber {
//			return array[l:r]
//		}
//		r = r - 1
//		TheNumber2 := TheNumber
//		TheLetter2 := TheLetter
//		index := 0
//		for i := r + 1; i < len(array); i++ {
//
//			if _, error := strconv.Atoi(array[i]); error == nil {
//				TheNumber2 = TheNumber2 + 1
//			} else {
//				TheLetter2 = TheLetter2 + 1
//			}
//
//			if _, error := strconv.Atoi(array[index]); error == nil {
//				TheNumber2 = TheNumber2 - 1
//			} else {
//				TheLetter2 = TheLetter2 - 1
//			}
//			index = index + 1
//			if TheLetter2 == TheNumber2 {
//				return array[index : i+1]
//			}
//
//		}
//
//	}
//
//	return []string{}
//
//}
