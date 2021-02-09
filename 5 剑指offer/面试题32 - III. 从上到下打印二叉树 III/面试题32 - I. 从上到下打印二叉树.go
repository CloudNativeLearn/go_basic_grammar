package main

func main() {

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	if root.Left == nil && root.Right == nil {
		return [][]int{{root.Val}}
	}
	Deque := []*TreeNode{}
	Result := [][]int{}
	Deque = append(Deque, root)
	ThisNeed := 1
	NextNeed := 0
	MindResult := []int{}
	Cheng := 1
	for len(Deque) != 0 {
		// 取出来一个
		mid := Deque[0]
		Deque = Deque[1:]
		ThisNeed = ThisNeed - 1
		//记录值

		if Cheng%2==0 {
			MindResult = append([]int{mid.Val}, MindResult...)
		}else {
			MindResult = append(MindResult, mid.Val)
		}

		// 把左边的放入


		if mid.Left != nil {
			Deque = append(Deque, mid.Left)
			NextNeed = NextNeed + 1
		}
		// 右边的放入
		if mid.Right != nil {
			Deque = append(Deque, mid.Right)
			NextNeed = NextNeed + 1
		}

		// 这一层循环完毕
		if ThisNeed == 0 {
			Result = append(Result, MindResult)
			NextNeed = 0
			ThisNeed = len(Deque)
			MindResult = []int{}
			Cheng = Cheng +1
		}
	}

	return Result

}
