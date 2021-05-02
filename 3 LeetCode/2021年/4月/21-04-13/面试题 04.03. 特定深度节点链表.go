package main

// 给定一棵二叉树，设计一个算法，创建含有某一深度上所有节点的链表（比如，若一棵树的深度为 D，则会创建出 D 个链表）。返回一个包含所有深度的链表的数组。
//
// 
//
//示例：
//
//输入：[1,2,3,4,5,null,7,8]
//
//        1
//       /  \
//      2    3
//     / \    \
//    4   5    7
//   /
//  8
//
//输出：[[1],[2,3],[4,5,7],[8]]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/list-of-depth-lcci
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type ListNode struct {
	Val  int
	Next *ListNode
}

func listOfDepth(tree *TreeNode) []*ListNode {
	if tree==nil{
		return nil
	}

	ThisCetnter := 0
	NextCenter := 0
	Result := []*ListNode{}
	MindNode := &ListNode{}


	addValueInListNode := func(listNode *ListNode,value int) {
		for listNode.Next!=nil{
			listNode = listNode.Next
		}
		listNode.Next = &ListNode{
			Val: value,
		}
	}

	Bfs := func(root *TreeNode) {
		AllNode := []*TreeNode{}
		AllNode = append(AllNode, root)
		ThisCetnter = 1
		for len(AllNode)!=0{
			x := AllNode[0]
			AllNode = AllNode[1:]
			if x.Left!=nil{
				AllNode = append(AllNode, x.Left)
				NextCenter = NextCenter+1
			}
			if x.Right!= nil{
				AllNode = append(AllNode, x.Right)
				NextCenter = NextCenter +1
			}
			ThisCetnter = ThisCetnter-1
			// 向中间数组
			addValueInListNode(MindNode,x.Val)
			// 如果一层遍历完
			if ThisCetnter==0{
				Result  = append(Result, MindNode.Next)
				ThisCetnter = NextCenter
				NextCenter = 0
				MindNode = &ListNode{}
			}

		}

	}

	Bfs(tree)
	return Result
}
