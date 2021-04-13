package main

import "fmt"

// 实现一个函数，检查一棵二叉树是否为二叉搜索树。
//
//示例 1:
//输入:
//    2
//   / \
//  1   3
//输出: true
//示例 2:
//输入:
//    5
//   / \
//  1   4
//     / \
//    3   6
//输出: false
//解释: 输入为: [5,1,4,null,null,3,6]。
//     根节点的值为 5 ，但是其右子节点值为 4 。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/legal-binary-search-tree-lcci
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	type zhongxu func(root *TreeNode)
	var zhong zhongxu
	list := []int{}
	zhong = func(root *TreeNode) {
		if root.Left != nil {
			zhong(root.Left)
		}
		list = append(list, root.Val)
		if root.Right != nil {
			zhong(root.Right)
		}
	}
	zhong(root)
	fmt.Println(list)
	panduan := func(arr []int) bool {
		if len(arr) <= 1 {
			return true
		}
		for i := 1; i < len(arr); i++ {
			if arr[i] <= arr[i-1] {
				return false
			}
		}
		return true

	}
	return panduan(list)
}
