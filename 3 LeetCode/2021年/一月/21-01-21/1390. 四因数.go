package main

import "fmt"

// 给你一个整数数组 nums，请你返回该数组中恰有四个因数的这些整数的各因数之和。
//如果数组中不存在满足题意的整数，则返回 0 。


//示例：
//输入：nums = [21,4,7]
//输出：32
//解释：
//21 有 4 个因数：阿里笔试, 3, 7, 21
//4 有 3 个因数：阿里笔试, 2, 4
//7 有 2 个因数：阿里笔试, 7
//答案仅为 21 的所有因数的和。

func main() {

	list := []int{21,4,7}
	fmt.Println(sumFourDivisors(list))

}


func sumFourDivisors(nums []int) int {
   result:=0
	for _,v:= range nums{
		result = result + TheForYIngNumber(v)
	}
	return  result
}

func TheForYIngNumber(number int) int  {
	result := 0
	TheIngNumber := 0
	leftMin := number
	for i:=1;i<leftMin;i++ {
		if number%i ==0 {
			if TheIngNumber<4 {
				// 是因数
				leftMin = number/i

				if i==leftMin {
					TheIngNumber = TheIngNumber +1
					result = result + i
				}else {
					TheIngNumber = TheIngNumber +2
					result = result + i+ leftMin
				}
			}else {
				return 0
			}
		}
	}
	if TheIngNumber!= 4 {
		return 0
	}else {
		return result
	}

}






