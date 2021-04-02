package main

import "fmt"

// 给你一个整数数组 nums，每次 操作 会从中选择一个元素并 将该元素的值减少 阿里笔试。
//
//如果符合下列情况之一，则数组 A 就是 锯齿数组：
//
//每个偶数索引对应的元素都大于相邻的元素，即 A[0] > A[阿里笔试] < A[2] > A[3] < A[4] > ...
//或者，每个奇数索引对应的元素都大于相邻的元素，即 A[0] < A[阿里笔试] > A[2] < A[3] > A[4] < ...
//返回将数组 nums 转换为锯齿数组所需的最小操作次数。
//
// 
//
//示例 阿里笔试：
//
//输入：nums = [阿里笔试,2,3]
//输出：2
//解释：我们可以把 2 递减到 0，或把 3 递减到 阿里笔试。
//示例 2：
//
//输入：nums = [9,6,阿里笔试,6,2]
//输出：4

func main() {
fmt.Println(movesToMakeZigzag([]int{10,4,4,10,10,6,2,3}))
}
//
//
// 遍历一次数组 同时记录偶数是大 和基数是大两种情况
func movesToMakeZigzag(nums []int) int {
	oddNumber := 0  //奇数
	evenNumber := 0 //偶数
	for index,_ := range nums{
	   if index%2 ==0{
	      // 偶数
          // 让自身 小于左右两边
          One := -1
          Thw := -1
          if index-1>=0 && nums[index]>=nums[index-1]{
               One = nums[index] - nums[index-1]+1
          }

          if index+1<len(nums) && nums[index]>=nums[index+1] {
               Thw = nums[index]-nums[index+1] +1
          }
          if One==-1&&Thw ==-1 {
             continue
          }else if One!=-1 && Thw==-1 {
             evenNumber = evenNumber + One
          }else if One==-1 && Thw!=-1 {
             evenNumber = evenNumber + Thw
          }else {

             if One>Thw {
                evenNumber = evenNumber+One
             }else {
                evenNumber = evenNumber+Thw
             }
          }
       }else {
          // 奇数
          // 让自身 小于左右两边
          One := -1
          Thw := -1
          if index-1>=0 && nums[index]>=nums[index-1]{
             One = nums[index] - nums[index-1]+1
          }

          if index+1<len(nums) && nums[index]>=nums[index+1] {
             Thw = nums[index]-nums[index+1] +1
          }

          if One==-1&&Thw ==-1 {
             continue
          }else if One!=-1 && Thw==-1 {
             oddNumber = oddNumber + One
          }else if One==-1 && Thw!=-1 {
             oddNumber = oddNumber + Thw
          }else {

             if One>Thw {
                oddNumber = oddNumber+One
             }else {
                oddNumber = oddNumber+Thw
             }
          }




       }
    }


   if oddNumber < evenNumber{
      return oddNumber
   }else {
      return evenNumber
   }

}
