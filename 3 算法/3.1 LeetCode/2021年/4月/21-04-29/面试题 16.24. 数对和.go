package main

//面试题 16.24. 数对和
//设计一个算法，找出数组中两数之和为指定值的所有整数对。一个数只能属于一个数对。
//
//示例 1:
//
//输入: nums = [5,6,5], target = 11
//输出: [[5,6]]
//示例 2:
//
//输入: nums = [5,6,5,6], target = 11
//输出: [[5,6],[5,6]]
//提示：
//
//nums.length <= 100000
//通过次数8,678提交次数18,743

func pairSums(nums []int, target int) [][]int {
	TheMap := map[int]int{}
	for _, v := range nums {
		TheMap[v] = TheMap[v] + 1
	}
	result := [][]int{}
	for k,v := range TheMap{
		if TheMap[target-k]!=0{
			if k==target-k{
				if v==1{
					TheMap[k]=0
					continue

				}
				y := v/2
				for i:=0;i<y;i++{
					result = append(result,[]int{k,k})
				}
				continue
			}
			if v>TheMap[target-k]{
				for i :=0;i<TheMap[target-k];i++{
					result = append(result, []int{k,target-k})
				}
				TheMap[k] = v - TheMap[target-k]
				TheMap[target-k] = 0
			}else {

				for i :=0;i<v;i++{
					result = append(result, []int{k,target-k})
				}
				TheMap[target-k] =  TheMap[target-k] -v
				TheMap[k] = 0
			}
		}
	}
	return result
}
