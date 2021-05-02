package main

import "fmt"

// 设计并实现一个算法，找出二叉树中某两个节点的第一个共同祖先。不得将其他的节点存储在另外的数据结构中。注意：这不一定是二叉搜索树。
//
//例如，给定如下二叉树: root = [3,5,阿里笔试,6,2,0,8,null,null,7,4]
//
//    3
//   / \
//  5   阿里笔试
// / \ / \
//6  2 0  8
//  / \
// 7   4
//示例 阿里笔试:
//
//输入: root = [3,5,阿里笔试,6,2,0,8,null,null,7,4], p = 5, q = 阿里笔试
//输出: 3
//解释: 节点 5 和节点 阿里笔试 的最近公共祖先是节点 3。
//示例 2:
//
//输入: root = [3,5,阿里笔试,6,2,0,8,null,null,7,4], p = 5, q = 4
//输出: 5
//解释: 节点 5 和节点 4 的最近公共祖先是节点 5。因为根据定义最近公共祖先节点可以为节点本身。
//说明:
//
//所有节点的值都是唯一的。
//p、q 为不同节点且均存在于给定的二叉树中。

func main() {
	Head := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 6,
				Left: nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 7,
					Left: nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val: 4,
					Left: nil,
					Right: nil,
				},
			},
		},
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 0,
				Left: nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val: 8,
				Left: nil,
				Right: nil,
			},
		},
	}
	p := &TreeNode{
		Val: 5,
	}
	q := &TreeNode{
		Val: 1,
	}

	fmt.Println(lowestCommonAncestor(Head,p,q))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 定义一个map 记录每个节点的父节点
var TheParent = map[int]*TreeNode{}
// 定义一个切片记录搜索路径
var visited = map[int]bool{}
func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
    dfs(root)
	for p!= nil {
		visited[p.Val] = true
		p = TheParent[p.Val]
	}

	for q!=nil  {
		if visited[q.Val]  {
			return q
		}
		q = TheParent[q.Val]
	}
	return nil
}

func dfs(root *TreeNode)  {
	if root == nil {
		return
	}
	if root.Left != nil{
		TheParent[root.Left.Val] = root
		dfs(root.Left)
	}
	if root.Right != nil{
		TheParent[root.Right.Val] = root
		dfs(root.Right)
	}
}