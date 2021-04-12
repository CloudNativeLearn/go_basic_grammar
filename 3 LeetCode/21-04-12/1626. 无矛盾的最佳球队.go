package main

import (
	"sort"
)
// 1626. 无矛盾的最佳球队
//假设你是球队的经理。对于即将到来的锦标赛，你想组合一支总体得分最高的球队。球队的得分是球队中所有球员的分数 总和 。
//
//然而，球队中的矛盾会限制球员的发挥，所以必须选出一支 没有矛盾 的球队。如果一名年龄较小球员的分数 严格大于 一名年龄较大的球员，则存在矛盾。同龄球员之间不会发生矛盾。
//
//给你两个列表 scores 和 ages，其中每组 scores[i] 和 ages[i] 表示第 i 名球员的分数和年龄。请你返回 所有可能的无矛盾球队中得分最高那支的分数 。
//
//
//
//示例 1：
//
//输入：scores = [1,3,5,10,15], ages = [1,2,3,4,5]
//输出：34
//解释：你可以选中所有球员。
//示例 2：
//
//输入：scores = [4,5,6,5], ages = [2,1,2,1]
//输出：16
//解释：最佳的选择是后 3 名球员。注意，你可以选中多个同龄球员。
//示例 3：
//
//输入：scores = [1,2,3,5], ages = [8,9,10,1]
//输出：6
//解释：最佳的选择是前 3 名球员。


func main() {
	bestTeamScore([]int{2,8,9},[]int{5,2,5})
}
//
//// 289  8 9 2
//// 525  2 5 5
//// sources 为分数  ages为年纪
//func bestTeamScore(scores []int, ages []int) int {
//
//	sources,agess := QuickSortStark(0,len(scores)-1,ages,scores)
//	arr := []int{}
//	for i:=0;i<len(scores);i++{
//		arr = append(arr,sources[i])
//		for j:=0;j<i;j++{
//			if sources[i]>=sources[j] {
//				arr[i] = GetMax2(arr[i],arr[j]+sources[i])
//			}else if agess[i]==agess[j]{
//				// 找到j 之前 不相等的agess
//				m := i
//				for m>=0{
//					if agess[m] != agess[i]{
//						break
//					}else{
//						m = m-1
//					}
//				}
//
//				if m==-1 || sources[i]>=sources[m]{
//					arr[i] = GetMax2(arr[i],arr[j]+sources[i])
//				}
//
//
//			}
//		}
//	}
//	fmt.Println(sources)
//	fmt.Println(agess)
//	fmt.Println(arr)
//	return GetMaxInArr(arr)
//}
//
//
//// 快速排序
//func QuickSortStark(l int,r int,ages []int,scores []int) ([]int ,[]int) {
//	x := (l+r)/2
//	Theleft := l
//	Theright := r
//	Xvalue := ages[x]
//	for l<r{
//		for ages[l]<Xvalue{
//			l = l+1
//		}
//		for  ages[r]>Xvalue {
//			r = r-1
//		}
//		if l>=r {
//			break
//		}
//		temp := ages[r]
//		ages[r] = ages[l]
//		ages[l] = temp
//
//		temp = scores[r]
//		scores[r] = scores[l]
//		scores[l] = temp
//
//		if ages[l] == Xvalue {
//			r = r-1
//		}
//		if ages[r] == Xvalue{
//			l = l+1
//		}
//	}
//
//	if l ==r {
//		l = l+1
//		r = r-1
//	}
//	if Theleft<l {
//		QuickSortStark(Theleft,r,ages,scores)
//	}
//	if Theright>r {
//		QuickSortStark(l,Theright,ages,scores)
//	}
//	return scores,ages
//}
//
//// 获得更大值
//func GetMax2(a int,b int)int{
//	if a>b{
//		return a
//	}
//	return b
//}
//
//// 获取数组中的最大值
//func GetMaxInArr(arr []int) int{
//	max := -1
//	for _,v := range arr{
//		max = GetMax2(v,max)
//	}
//	return max
//}


func bestTeamScore(scores []int, ages []int) int {
	l:= len(scores)
	max := 0
	arr:=make([][2]int,l)
	for i,v :=range ages {
		arr[i] = [2]int{v, scores[i]}
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] < arr[j][0] || (arr[i][0] == arr[j][0] && arr[i][1] < arr[j][1])
	})
	dp:=make([]int, l)
	for i := 0; i < l; i++ {
		for j := i - 1; j >= 0; j-- {
			if (arr[j][1] <= arr[i][1]) {
				if dp[i] < dp[j]{
					dp[i] = dp[j]
				}
			}
		}
		dp[i] += arr[i][1]
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}

