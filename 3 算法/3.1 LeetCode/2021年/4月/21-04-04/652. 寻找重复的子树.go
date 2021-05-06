package main

import (
	"fmt"
	"strconv"
)

// 给定一棵二叉树，返回所有重复的子树。对于同一类的重复子树，你只需要返回其中任意一棵的根结点即可。
//
//两棵树重复是指它们具有相同的结构以及相同的结点值。
//
//示例 1：
//
//        1
//       / \
//      2   3
//     /   / \
//    4   2   4
//       /
//      4
//下面是两个重复的子树：
//
//      2
//     /
//    4
//和
//
//    4
//因此，你需要以列表的形式返回上述重复子树的根结点。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/find-duplicate-subtrees
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

func main() {
	//root := &TreeNode{
	//	Val: 1,
	//	Left: &TreeNode{
	//		Val: 2,
	//		Left: &TreeNode{
	//			Val: 4,
	//		},
	//	},
	//	Right: &TreeNode{
	//		Val: 3,
	//		Left: &TreeNode{
	//			Val: 2,
	//			Left: &TreeNode{
	//				Val: 4,
	//			},
	//		},
	//		Right: &TreeNode{
	//			Val: 4,
	//		},
	//	},
	//}

	//root:=&TreeNode{
	//	Val: 2,
	//	Left: &TreeNode{
	//		Val: 1,
	//	},
	//	Right: &TreeNode{
	//		Val: 1,
	//	},
	//}


	//root:= &TreeNode{
	//	Val: 0,
	//	Left: &TreeNode{
	//		Val: 0,
	//		Left: &TreeNode{
	//			Val: 0,
	//		},
	//	},
	//	Right: &TreeNode{
	//		Val: 0,
	//		Right: &TreeNode{
	//			Val: 0,
	//			Right: &TreeNode{
	//				Val: 0,
	//			},
	//		},
	//	},
	//}

	root := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 11,
			},
		},
		Right: &TreeNode{
			Val: 11,
			Left: &TreeNode{
				Val: 1,
			},
		},
	}
	fmt.Println(findDuplicateSubtrees(root))
}


func findDuplicateSubtrees(root *TreeNode) []*TreeNode {

	HashMap := map[string]int{}
	arr := []*TreeNode{}
	Dfs(root, &HashMap, &arr)
	return arr
}

// 深度优先遍历数

func Dfs(root *TreeNode, HashMap *map[string]int, arr *[]*TreeNode) string {
	str2 := strconv.Itoa(root.Val)
	str2 = str2+" "
	if root.Right == nil && root.Left == nil {
		str2 = strconv.Itoa(root.Val)
		str2 = str2 + " "
	}
	strL := " null "
	if root.Left != nil {
		strL = Dfs(root.Left, HashMap, arr)
	}
	strR := " null "
	if root.Right != nil {
		strR = Dfs(root.Right, HashMap, arr)
	}
	if (*HashMap)[str2+strL+strR] ==0 {
		(*HashMap)[str2+strL+strR] =1
	}else if (*HashMap)[str2+strL+strR]==1 {
		(*HashMap)[str2+strL+strR]=2
		(*arr) = append(*arr, root)
	}else {
		(*HashMap)[str2+strL+strR] = (*HashMap)[str2+strL+strR]+1
	}

	return  str2+strL + strR
}
