package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(summaryRanges([]int{0,1,2,4,5,7}))
}



func summaryRanges(nums []int) []string {
   s := []string{}

   l :=0
   r :=0
	for  {
		if r+1<len(nums){
			if nums[r+1] == nums[r]+1 {
				// 连续
				r = r +1
				continue
			}else {
				// 不连续 将现有区间加入s
				// 判断 l和r是否相等 如果相等 只添加数字
				if len(nums)==0{
					return []string{}
				}
				if l == r {
					s = append(s,strconv.Itoa(nums[r]))
				}else {
					str :=  strconv.Itoa(nums[l]) + "->" + strconv.Itoa(nums[r])
					s =  append(s, str)
				}
				l = r+1
				r = r+1
			}
		}else {
			// 已经到了尾部
			// 不连续 将现有区间加入s
			// 判断 l和r是否相等 如果相等 只添加数字
			if l == r {
				s = append(s,strconv.Itoa(nums[r]))
			}else {
				str :=  strconv.Itoa(nums[l]) + "->" + strconv.Itoa(nums[r])
				s =  append(s, str)
			}

			return s
		}

	}

}
