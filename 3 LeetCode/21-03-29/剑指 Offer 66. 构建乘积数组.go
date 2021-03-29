package main

// 给定一个数组 A[0,1,…,n-1]，请构建一个数组 B[0,1,…,n-1]，其中 B[i] 的值是数组 A 中除了下标 i 以外的元素的积, 即 B[i]=A[0]×A[1]×…×A[i-1]×A[i+1]×…×A[n-1]。不能使用除法。
//
// 
//
//示例:
//
//输入: [1,2,3,4,5]
//输出: [120,60,40,30,24]
// 
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/gou-jian-cheng-ji-shu-zu-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//func main() {
//fmt.Println(constructArr([]int{1,2,3,4}))
//}

func constructArr(a []int) []int {
	if len(a)==0 {
		return []int{}
	}
	//  构建一个从左乘到右的数组
	left := []int{}
	//  构建一个从右乘到左的数组
	right := []int{}
	for i,v := range a{
		if i-1<0 {
			left = append(left, v)
			right = append(right,a[len(a)-1])
		}else {
			left = append(left, v*left[i-1])
			right = append(right, right[i-1]*a[len(a)-i-1])
		}
	}
	result := []int{}
	result = append(result, right[len(a)-1-1])
	for i:=1;i<len(a)-1;i++{
		result = append(result, left[i-1]*right[len(a)-i-1-1])
	}
	result = append(result, left[len(a)-1-1])
	return result


}
