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

func levelOrder(root *TreeNode) []int {
	if root==nil {
		return []int{}
	}
	Result :=[]int{}
	Deque := []*TreeNode{}
	Deque = append(Deque,root)
	for len(Deque) !=0 {
		mid := Deque[0]
		Deque = Deque[1:]
		Result = append(Result,mid.Val)
		if mid.Left!=nil {
			Deque = append(Deque,mid.Left)
		}
		if mid.Right!=nil {
			Deque = append(Deque,mid.Right)
		}
	}
	return Result

}
