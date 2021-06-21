package main

// 根据一棵树的前序遍历与中序遍历构造二叉树。
//
//注意:
//你可以假设树中没有重复的元素。
//
//例如，给出
//
//前序遍历 preorder =[3,9,20,15,7]
//中序遍历 inorder = [9,3,15,20,7]
//返回如下的二叉树：
//
//    3
//   / \
//  9  20
//    /  \
//   15   7
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//
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
	buildTree([]int{3,9,20,15,7},[]int{9,3,15,20,7})
}

//前序遍历 preorder =[3,9,20,15,7]
//中序遍历 inorder = [9,3,15,20,7]

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 1 {
		return &TreeNode{
			Val: preorder[0],
		}
	}else if len(preorder)==0{
		return nil
	}

	// 根节点
	root := preorder[0]
	// 在中序中的节点
	index_inInorder := GetIndex(root, inorder)

	// 左节点
	left_preorder :=preorder[1:index_inInorder+1]
	left_inorder  :=inorder[0:index_inInorder]
	// 右节点
	right_preorder :=preorder[index_inInorder+1:]
	right_inorder  :=inorder[index_inInorder+1:]

	result := &TreeNode{Val: root}
	result.Left = buildTree(left_preorder,left_inorder)
	result.Right = buildTree(right_preorder,right_inorder)
	return result


}

func GetIndex(a int, arr []int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == a {
			return i
		}
	}
	return 0
}
