package main

// 幂集。编写一种方法，返回某集合的所有子集。集合中不包含重复的元素。
//
//说明：解集不能包含重复的子集。
//
//示例:
//
// 输入： nums = [1,2,3]
// 输出：
//[
//  [3],
//  [1],
//  [2],
//  [1,2,3],
//  [1,3],
//  [2,3],
//  [1,2],
//  []
//]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/power-set-lcci
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	subsets([]int{0,1,2,3,4})
}

func subsets(nums []int) [][]int {
	result := [][]int{}
	result = append(result, []int{})
	for i := 0; i < len(nums); i++ {
		x := len(result)
		newArr := [][]int{}
		for j:=0;j<x;j++{
			y := make([]int,len(result[j]))
			copy(y,result[j])
			y = append(y, nums[i])
			newArr = append(newArr, y)
		}
		result = append(result, newArr...)
	}
	return result
}
