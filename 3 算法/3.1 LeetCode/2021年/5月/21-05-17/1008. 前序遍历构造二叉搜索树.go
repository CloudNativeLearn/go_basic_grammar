package main

import "fmt"

// 返回与给定前序遍历 preorder 相匹配的二叉搜索树（binary search tree）的根结点。
//
//(回想一下，二叉搜索树是二叉树的一种，其每个节点都满足以下规则，对于 node.left 的任何后代，值总 < node.val，而 node.right 的任何后代，值总 > node.val。此外，前序遍历首先显示节点 node 的值，然后遍历 node.left，接着遍历 node.right。）
//
//题目保证，对于给定的测试用例，总能找到满足要求的二叉搜索树。
//
// 
//
//示例：
//
//输入：[8,5,1,7,10,12]
//输出：[8,5,10,1,7,null,12]
//
// 
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/construct-binary-search-tree-from-preorder-traversal
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

func bstFromPreorder(preorder []int) *TreeNode {
	inorder:= make([]int,len(preorder))
	copy(inorder,preorder)
	QuickSortStark(0,len(inorder)-1,inorder)
	return buildTree(preorder,inorder)

}


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


func QuickSortStark(l int,r int,arr []int)  {
	x := (l+r)/2
	Theleft := l
	Theright := r
	Xvalue := arr[x]
	for l<r{

		// 在右边找到比 xvalue小的值
		for arr[l]<Xvalue{
			l = l+1
		}

		// 在左边找到比 xvalue大的值
		for  arr[r]>Xvalue {
			r = r-1
		}

		if l>=r {
			break
		}

		temp := arr[r]
		arr[r] = arr[l]
		arr[l] = temp

		if arr[l] == Xvalue {
			r = r-1
		}
		if arr[r] == Xvalue{
			l = l+1
		}

	}

	if l ==r {
		l = l+1
		r = r-1
	}

	if Theleft<l {
		QuickSortStark(Theleft,r,arr)
	}
	if Theright>r {
		QuickSortStark(l,Theright,arr)
	}
}

func GetIndex(a int, arr []int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == a {
			return i
		}
	}
	return 0
}
