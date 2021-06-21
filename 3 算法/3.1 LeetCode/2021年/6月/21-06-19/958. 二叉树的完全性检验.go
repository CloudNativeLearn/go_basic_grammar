package main

// 给定一个二叉树，确定它是否是一个完全二叉树。
//
//百度百科中对完全二叉树的定义如下：
//
//若设二叉树的深度为 h，除第 h 层外，其它各层 (1～h-1) 的结点数都达到最大个数，第 h 层所有的结点都连续集中在最左边，这就是完全二叉树。（注：第 h 层可能包含 1~ 2h 个节点。）
//
// 
//
//示例 1：
//
//
//
//输入：[1,2,3,4,5,6]
//输出：true
//解释：最后一层前的每一层都是满的（即，结点值为 {1} 和 {2,3} 的两层），且最后一层中的所有结点（{4,5,6}）都尽可能地向左。
//示例 2：
//
//
//
//输入：[1,2,3,4,5,null,7]
//输出：false
//解释：值为 7 的结点没有尽可能靠向左侧。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/check-completeness-of-a-binary-tree
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//
//func main() {
//	root := TreeNode{
//		Val: 1,
//		Left: &TreeNode{
//			Val: 2,
//			Left: &TreeNode{
//				Val: 4,
//			},
//			Right: &TreeNode{
//				Val: 5,
//			},
//		},
//		Right: &TreeNode{
//			Val: 3,
//			Left: &TreeNode{
//				Val: 6,
//			},
//		},
//	}
//
//	isCompleteTree(&root)
//}
//* Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
//
//func isCompleteTree(root *TreeNode) bool {
//	list := []*TreeNode{}
//	selfNumber := 1
//	nextNumber := 0
//	list = append(list, root)
//
//
//	for len(list) != 0 {
//		one := list[0]
//		//fmt.Println(one.Val)
//		list = list[1:]
//
//		if one.Left != nil {
//			list = append(list, one.Left)
//			nextNumber = nextNumber + 1
//		}
//
//		if one.Right != nil {
//			list = append(list, one.Right)
//			nextNumber = nextNumber + 1
//		}
//		selfNumber = selfNumber - 1
//		if selfNumber == 0 {
//			//fmt.Println("一层遍历结束")
//			selfNumber = nextNumber
//			nextNumber = 0
//		}
//
//	}
//
//	return true
//}


func isCompleteTree(root *TreeNode) bool {
	end := false
	queue := []*TreeNode{root}
	for len(queue)>0{
		one := queue[0]
		queue = queue[1:]
		if one ==nil{
			end = true
		}else {
			if end == true{
				return false
			}
			queue = append(queue, one.Left)
			queue = append(queue, one.Right)

		}



	}

	return true
}