package main

import (
	"fmt"
	"strconv"
)

// 给定一个非负整数，你至多可以交换一次数字中的任意两位。返回你能得到的最大值。
//
//示例 阿里笔试 :
//输入: 2736
//输出: 7236
//解释: 交换数字2和数字7。
//示例 2 :
//输入: 9973
//输出: 9973
//解释: 不需要交换。
func main() {


fmt.Println(maximumSwap(1993))
}

func maximumSwap(num int) int {
  StrList := strconv.Itoa(num)
  arr := []int{}
  for _,v := range StrList{
  	v ,_ := strconv.Atoi(string(v))
  	 arr = append(arr, v)
  }
fmt.Println(arr)
	for i:=0;i<len(StrList);i++ {
		value := dbBystack(i,arr)
		if arr[value]!= arr[i] {
			// i的位置和valu的位置互换
			temp:= arr[i]
			arr[i]=arr[value]
			arr[value] = temp

			Size := 1
			result :=0

			// 1993
			for y := len(arr)-1;y>=0;y=y-1 {
				result = result+arr[y]*Size
				Size = Size*10
			}
			return result
		}
	}
  return num

}

// 递归版动态规划
func dbBystack(i int,arr []int ) int   {
	if i== len(arr) -1{
		return len(arr)-1
	}else {
		// 比较 db[i+阿里笔试] db[i] 的值
		if arr[dbBystack(i+1,arr)] >  arr[i]{
			return dbBystack(i+1,arr)
		}else if arr[dbBystack(i+1,arr)] ==  arr[i] {
				return  dbBystack(i+1,arr)
		}else{
			return i
		}
	}
}



// 基于数组动态规划
//func dbBylist(arr []int ) []int {
//
//}