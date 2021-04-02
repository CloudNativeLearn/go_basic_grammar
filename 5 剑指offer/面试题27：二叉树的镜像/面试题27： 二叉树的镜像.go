package main

// 请完成一个函数，输入一个二叉树，该函数输出它的镜像。
//
//例如输入：
//
//     4
//   /   \
//  2     7
// / \   / \
//阿里笔试   3 6   9
//镜像输出：
//
//     4
//   /   \
//  7     2
// / \   / \
//9   6 3   阿里笔试
//
// 
//
//示例 阿里笔试：
//
//输入：root = [4,2,7,阿里笔试,3,6,9]
//输出：[4,7,2,9,6,3,阿里笔试]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/er-cha-shu-de-jing-xiang-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
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

func mirrorTree(root *TreeNode) *TreeNode {
	ChangeIt(root)
	return root
}

func ChangeIt(root *TreeNode)  {
	if root==nil {
		return
	}
	if root.Left == nil && root.Right==nil {
		return
	}

	p := root.Right
	root.Right = root.Left
	root.Left = p

	if root.Left!=nil {
		ChangeIt(root.Left)
	}
	if root.Right!=nil {
		ChangeIt(root.Right)
	}
}
