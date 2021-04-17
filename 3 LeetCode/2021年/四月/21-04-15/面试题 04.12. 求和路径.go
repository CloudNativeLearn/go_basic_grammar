package main

// 给定一棵二叉树，其中每个节点都含有一个整数数值(该值或正或负)。设计一个算法，打印节点数值总和等于某个给定值的所有路径的数量。注意，路径不一定非得从二叉树的根节点或叶节点开始或结束，但是其方向必须向下(只能从父节点指向子节点方向)。
//
//示例:
//给定如下二叉树，以及目标和 sum = 22，
//
//              5
//             / \
//            4   8
//           /   / \
//          11  13  4
//         /  \    / \
//        7    2  5   1
//返回:
//
//3
//解释：和为 22 的路径有：[5,4,11,2], [5,8,4,5], [4,11,7]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/paths-with-sum-lcci
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

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



func pathSum(root *TreeNode, sum int) int {

	if root==nil{
		return 0
	}
	value:= 0
	value = Stark(root,0,sum)
	left := 0
	if root.Left!=nil{
		left = pathSum(root.Left,sum)
	}

	right :=0
	if root.Right!=nil{
		right = pathSum(root.Right,sum)
	}



	return value+right+left

}

func Stark(root *TreeNode,mindresult int,sum int) int  {
	if root==nil{
		return 0
	}
	left := 0
	if root.Left!=nil{
		left = Stark(root.Left,mindresult+root.Val,sum)
	}
	right := 0
	if root.Right!=nil{
		right = Stark(root.Right,mindresult+root.Val,sum)
	}
	if mindresult+root.Val==sum{
		return left+right +1
	}else {
		return left+right
	}
}