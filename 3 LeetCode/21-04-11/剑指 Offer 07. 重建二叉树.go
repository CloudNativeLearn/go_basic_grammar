package main

// 输入某二叉树的前序遍历和中序遍历的结果，请重建该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
//
// 
//
//例如，给出
//
//前序遍历 preorder = [3,9,20,15,7]
//中序遍历 inorder = [9,3,15,20,7]
//返回如下的二叉树：
//
//    3
//   / \
//  9  20
//    /  \
//   15   7
// 
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/zhong-jian-er-cha-shu-lcof
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


// preorder为前序遍历  inorder 为后续遍历
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder)==0||  len(inorder)==0 {
		return nil
	}
	if len(preorder) != len(inorder) {
		return nil
	}

	root := &TreeNode{Val: preorder[0] } // root 节点为前序中的第一个
	pos := FindInInorder(root.Val,inorder)
	preorder_left := preorder[1:pos+1]
	preorder_right := preorder[pos+1:]

	inorder_left := inorder[:pos]
	inorder_right := inorder[pos+1:]

	left := buildTree(preorder_left,inorder_left)
	right := buildTree(preorder_right,inorder_right)
	root.Left = left
	root.Right = right

	return root



}

// 在中序节点中找到对应的值
func FindInInorder(b int,arr []int)  int{
	for i:=0;i<len(arr);i++{
		if arr[i] ==b{
			return i
		}
	}
	return -1
}




