package main

import "fmt"

// 给你一个二叉树的根节点 root。设根节点位于二叉树的第 1 层，而根节点的子节点位于第 2 层，依此类推。
//
//请你找出层内元素之和 最大 的那几层（可能只有一层）的层号，并返回其中 最小 的那个。

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

func main() {
root := &TreeNode{
	Val: 1,
	Left: &TreeNode{
		Val: 7,
		Left: &TreeNode{
			Val: 7,
			Left: nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: -8,
			Left: nil,
			Right: nil,
		},
	},
	Right: &TreeNode{
		Val: 0,
		Left: nil,
		Right: nil,
	},
}
println(maxLevelSum(root))
}


func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	Stack:=[]*TreeNode{}
	Stack = append(Stack, root)
	result := []int{}
	value :=0
	ThisNeed := 1
	TheNextNeed :=0
	for len(Stack)!=0 {
		// 取出stack第一个数字
		thefirst := GetFirst(&Stack)
		Stack = Stack[1:]
		value = value + thefirst.Val
		ThisNeed = ThisNeed-1

		if thefirst.Left!=nil {
			Stack = append(Stack, thefirst.Left)
			TheNextNeed=TheNextNeed+1
		}
		if thefirst.Right != nil{
			Stack = append(Stack,thefirst.Right)
			TheNextNeed = TheNextNeed+1
		}
		// 一层遍历完成
		if ThisNeed==0 {
			result = append(result, value)
			ThisNeed =TheNextNeed
			TheNextNeed = 0
			value = 0
		}


	}

	TheMaxIndex :=0
	TheMaxValue :=result[0]
	for k,v := range result{
		if TheMaxValue<v {
			TheMaxValue = v
			TheMaxIndex = k
		}
	}
	fmt.Println(result)
	return TheMaxIndex+1
}



// 取出第一个数据
func GetFirst(arr *[]*TreeNode) *TreeNode{
  first := (*arr)[0]
  ThNewArr := (*arr)[1:]
  arr = &ThNewArr
  return first
}
