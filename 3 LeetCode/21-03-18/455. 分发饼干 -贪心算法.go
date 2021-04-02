package main

import "fmt"

// 假设你是一位很棒的家长，想要给你的孩子们一些小饼干。但是，每个孩子最多只能给一块饼干。
//
//对每个孩子 i，都有一个胃口值 g[i]，这是能让孩子们满足胃口的饼干的最小尺寸；并且每块饼干 j，都有一个尺寸 s[j] 。如果 s[j] >= g[i]，我们可以将这个饼干 j 分配给孩子 i ，这个孩子会得到满足。你的目标是尽可能满足越多数量的孩子，并输出这个最大数值。
//
// 
//示例 阿里笔试:
//
//输入: g = [阿里笔试,2,3], s = [阿里笔试,阿里笔试]
//输出: 阿里笔试
//解释:
//你有三个孩子和两块小饼干，3个孩子的胃口值分别是：阿里笔试,2,3。
//虽然你有两块小饼干，由于他们的尺寸都是1，你只能让胃口值是1的孩子满足。
//所以你应该输出1。
//示例 2:
//
//输入: g = [阿里笔试,2], s = [阿里笔试,2,3]
//输出: 2
//解释:
//你有两个孩子和三块小饼干，2个孩子的胃口值分别是1,2。
//你拥有的饼干数量和尺寸都足以让所有孩子满足。
//所以你应该输出2.

func main() {
fmt.Println(findContentChildren([]int{1,2},[]int{3,2,1}))
}


func findContentChildren(g []int, s []int) int {
// 把g和s排序
	g = XierSort(g)
	s = XierSort(s)

	result := 0
		g0 := 0
		s0 := 0
	if len(g) ==0 || len(s)==0 {
		return result
	}
	for  {
		if g0==len(g) || s0==len(s) {
			return result
		}else if g[g0]<=s[s0]  {
			result = result+1
			g0 = g0+1
			s0=s0+1
			continue
		}else if s[s0]<g[g0] {
			s0 = s0+1
			continue
		}
	}


}


// 希尔排序  从小到大
func XierSort(arr []int)[]int  {
	number := len(arr)
	for number !=1 && number!=0{
		number = number/2
		for i:=number;i<len(arr);i++{
			for j:=i-number;j>=0;j=j-number{
				if arr[j]>arr[j+number] {
					// 交换两个数
					temp := arr[j]
					arr[j] = arr[j+number]
					arr[j+number] = temp
				}
			}
		}
	}
	return arr
}