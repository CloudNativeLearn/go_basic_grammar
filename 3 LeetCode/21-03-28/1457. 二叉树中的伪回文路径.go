package main

import "fmt"

// 给你一棵二叉树，每个节点的值为 1 到 9 。我们称二叉树中的一条路径是 「伪回文」的，当它满足：路径经过的所有节点值的排列中，存在一个回文序列。
//
//请你返回从根到叶子节点的所有路径中 伪回文 路径的数目。
//
// 
//
//示例 1：
//
//
//
//输入：root = [2,3,1,3,1,null,1]
//输出：2
//解释：上图为给定的二叉树。总共有 3 条从根到叶子的路径：红色路径 [2,3,3] ，绿色路径 [2,1,1] 和路径 [2,3,1] 。
//     在这些路径中，只有红色和绿色的路径是伪回文路径，因为红色路径 [2,3,3] 存在回文排列 [3,2,3] ，绿色路径 [2,1,1] 存在回文排列 [1,2,1] 。
//示例 2：
//
//
//
//输入：root = [2,1,1,1,3,null,null,null,null,null,1]
//输出：1
//解释：上图为给定二叉树。总共有 3 条从根到叶子的路径：绿色路径 [2,1,1] ，路径 [2,1,3,1] 和路径 [2,1] 。
//     这些路径中只有绿色路径是伪回文路径，因为 [2,1,1] 存在回文排列 [1,2,1] 。
//示例 3：
//
//输入：root = [9]
//输出：1
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/pseudo-palindromic-paths-in-a-binary-tree
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	root := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 1,
			Left: nil,
			Right: &TreeNode{
				Val: 1,
			},
		},
	}

	fmt.Println(pseudoPalindromicPaths(root))
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

func pseudoPalindromicPaths(root *TreeNode) int {
	MindResult := &map[int]int{}
	Number := 0
	Defs(root, MindResult,&Number)
	return Number
}

// 深度遍历二叉树 查看所有路径
func Defs(root *TreeNode, Map *map[int]int,Number *int) {

	(*Map)[root.Val] = (*Map)[root.Val] +1
	if root.Right == nil && root.Left == nil {
		// 左右都为 nill 向result中添加
		if DealList(Map){
			*Number = *Number +1

		}

	}
	if root.Right != nil {
		// 中间结果数组 拷贝一份
		Mid := DeepCopyMap(*Map)
		Defs(root.Right, &Mid,Number)
	}
	if root.Left != nil {
		// 中间结果数组 拷贝一份
		Mid := DeepCopyMap(*Map)
		Defs(root.Left, &Mid,Number)
	}
}

// 循环遍历二维数组 统计每个数组中的数字
func DealList(arr *map[int]int) bool {

	Ji := 0
	//fmt.Println(NewMap)
	for _,v := range *arr{
		if v%2 !=0{
			Ji = Ji+1
		}
	}
	//fmt.Println(Ji)
	//fmt.Println("=================")
	return Ji==1 || Ji==0

}



// map的深度拷贝
func DeepCopyMap(mp map[int]int)map[int]int  {
	newMap := map[int]int{}
	for k,v := range mp{
		newMap[k] = v
	}
	return newMap
}