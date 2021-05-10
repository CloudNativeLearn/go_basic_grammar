package main

// 给定一个整数数组 A，坡是元组 (i, j)，其中  i < j 且 A[i] <= A[j]。这样的坡的宽度为 j - i。
//
//找出 A 中的坡的最大宽度，如果不存在，返回 0 。
//
// 
//
//示例 1：
//
//输入：[6,0,8,2,1,5]
//输出：4
//解释：
//最大宽度的坡为 (i, j) = (1, 5): A[1] = 0 且 A[5] = 5.
//示例 2：
//
//输入：[9,8,1,0,1,9,4,0,4,1]
//输出：7
//解释：
//最大宽度的坡为 (i, j) = (2, 9): A[2] = 1 且 A[9] = 1.
// 
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/maximum-width-ramp
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func maxWidthRamp(A []int) int {
	if len(A)==0{
		return 0
	}
	if len(A)==1{
		return 1
	}
	stark := Stark{Arr: []int{}}
	for k, v := range A {
		if stark.IsEmpty() {
			stark.Push(k)
		} else if A[stark.GetEnd()] > v {
			// fmt.Println(A[stark.GetEnd()],v)
			stark.Push(k)
		}
	}
	// fmt.Println(stark.Arr)
	result := 0
	for i := len(A) - 1; i >= 0; i-- {
		for !stark.IsEmpty() && A[stark.GetEnd()] <= A[i] {

			result = GetMax(result,i-stark.Pop())

		}
	}
	return result

}

func AXIaoB(a int, b int) bool {
	if a < b {
		return true
	} else {
		return false
	}
}

func GetMax(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

type Stark struct {
	Arr []int
}

func (this *Stark) Push(a int) {
	this.Arr = append(this.Arr, a)
}

func (this *Stark) Pop() int {
	a := this.Arr[len(this.Arr)-1]
	this.Arr = this.Arr[:len(this.Arr)-1]
	return a
}

func (this *Stark) IsEmpty() bool {
	if len(this.Arr) == 0 {
		return true
	}
	return false
}

func (this *Stark) GetEnd() int {
	return this.Arr[len(this.Arr)-1]
}