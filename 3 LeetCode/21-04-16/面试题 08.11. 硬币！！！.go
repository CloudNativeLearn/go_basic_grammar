package main

// 硬币。给定数量不限的硬币，币值为25分、10分、5分和1分，编写代码计算n分有几种表示法。(结果可能会很大，你需要将结果模上1000000007)
//
//示例1:
//
// 输入: n = 5
// 输出：2
// 解释: 有两种方式可以凑成总金额:
//5=5
//5=1+1+1+1+1
//示例2:
//
// 输入: n = 10
// 输出：4
// 解释: 有四种方式可以凑成总金额:
//10=10
//10=5+5
//10=5+1+1+1+1+1
//10=1+1+1+1+1+1+1+1+1+1
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/coin-lcci
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//func main() {
//	fmt.Println(waysToChange(10))
//}
//func waysToChange(n int) int {
//	result := make([]int, n+1)
//	consume := []int{1, 5, 10, 25}
//	result[0] = 1
//	for k,_ := range result{
//		for _,cV := range consume{
//			if k-cV>=0{
//				result[k] = result[k] + result[k-cV]
//			}
//		}
//	}
//	return result[len(result)-1]%1000000007
//
//}
