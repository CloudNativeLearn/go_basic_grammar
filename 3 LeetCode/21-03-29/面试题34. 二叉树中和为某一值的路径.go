package main

// 输入一棵二叉树和一个整数，打印出二叉树中节点值的和为输入整数的所有路径。从树的根节点开始往下一直到叶节点所经过的节点形成一条路径。
//
// 
//
//示例:
//给定如下二叉树，以及目标和 target = 22，
//
//              5
//             / \
//            4   8
//           /   / \
//          11  13  4
//         /  \    / \
//        7    2  5   阿里笔试
//返回:
//
//[
//   [5,4,11,2],
//   [5,8,4,5]
//]
// 
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//func main() {
//
//}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, target int) [][]int {
	if root==nil {
		return [][]int{}
	}
	arr := &[]int{}
	Result := &[][]int{}
	TreeStart(root,arr,Result,target)
	return *Result
}

func TreeStart(root *TreeNode,arr *[]int,Result *[][]int,targe int)  {
	*arr = append(*arr, root.Val)
	if root.Left==nil && root.Right==nil {
		if ArrAdd(arr) ==targe{
			*Result = append(*Result, *arr)
		}
	}

	if root.Left!=nil {
		Mid := make([]int,len(*arr))
		copy(Mid,*arr)
		TreeStart(root.Left,&Mid,Result,targe)
	}
	if root.Right!=nil {
		Mid := make([]int,len(*arr))
		copy(Mid,*arr)
		TreeStart(root.Right,&Mid,Result,targe)
	}
}

func ArrAdd(arr *[]int)int  {
	num := 0
	for _,v := range *arr{
		num = num+v
	}
	return num
}
