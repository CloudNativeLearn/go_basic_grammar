package main

// 输入两个整数序列，第一个序列表示栈的压入顺序，请判断第二个序列是否为该栈的弹出顺序。假设压入栈的所有数字均不相等。例如，序列 {1,2,3,4,5} 是某栈的压栈序列，序列 {4,5,3,2,1} 是该压栈序列对应的一个弹出序列，但 {4,3,5,1,2} 就不可能是该压栈序列的弹出序列。
//
// 
//
//示例 1：
//
//输入：pushed = [1,2,3,4,5], popped = [4,5,3,2,1]
//输出：true
//解释：我们可以按以下顺序执行：
//push(1), push(2), push(3), push(4), pop() -> 4,
//push(5), pop() -> 5, pop() -> 3, pop() -> 2, pop() -> 1
//示例 2：
//
//输入：pushed = [1,2,3,4,5], popped = [4,3,5,1,2]
//输出：false
//解释：1 不能在 2 之前弹出。

func main() {

	list1 := []int{1,0}
	list2 := []int{1,0}
	validateStackSequences(list1,list2)
}

func validateStackSequences(pushed []int, popped []int) bool {
	if len(pushed)==0 {
		return true
	}
	left :=0
	right:=0
	for len(pushed)!=0 {

		if left>=len(pushed) {
			return false
		}
		if len(pushed)==0 {
			return true
		}
		if pushed[left]==popped[right] {
			pushed = append(pushed[:left], pushed[left+1:]...)
			if left!=0 {
				left = left-1
			}
			right = right +1
			continue
		}

		left = left +1

	}
	return true
}