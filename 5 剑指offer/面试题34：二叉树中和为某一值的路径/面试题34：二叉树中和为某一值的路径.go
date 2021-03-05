package main

// 输入一棵二叉树和一个整数，打印出二叉树中节点值的和为输入整数的所有路径。从树的根节点开始往下一直到叶节点所经过的节点形成一条路径。
//
// 
//
//示例:
//给定如下二叉树，以及目标和 sum = 22，
//
//              5
//             / \
//            4   8
//           /   / \
//          11  13  4
//         /  \    / \
//        7    2  5   1
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
func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	list := []*TreeNode{}
	result := [][]int{}
	Talk(list,root,sum,&result)
 return result
}

func Talk(list []*TreeNode,root *TreeNode,Sum int,resultList *[][]int)  {
	// 判断root
	if(root==nil){
		return
	}else if root.Val>Sum {
		return
	}else if (root.Val<Sum){
		list=append(list,root)
		Sum = Sum-root.Val
	}

	// 判断左边
	//     小于sum 直接进入然后递归
	//	   大于sum 直接判断右边
	//	   如果等于nil  判断和  判断右边
	if(root.Left==nil){
		if Sum==0 {
			OneList := []int{}
			for _,v := range list{
				OneList = append(OneList, v.Val)
			}
			*resultList = append(*resultList,OneList)
		}

	}else if root.Left.Val>Sum {

	}else if (root.Left.Val<Sum){
		Talk(list,root.Left,Sum,resultList)
	}


	// 判断右边
	//	   小于sum 直接进入然后递归
	//	   大于sum  将list中的路径最后一个删除  sum加上该节点的值
	//     等于nill  判断和 将list中的路径最后一个删除  sum加上该节点的值

	if(root.Right==nil){
		if Sum==0 {
			OneList := []int{}
			for _,v := range list{
				OneList = append(OneList, v.Val)
			}
			*resultList = append(*resultList,OneList)
		}

		Sum = Sum+root.Val
		list = list[:len(list)-1]

	}else if root.Left.Val>Sum {
		if Sum==0 {
			OneList := []int{}
			for _,v := range list{
				OneList = append(OneList, v.Val)
			}
			*resultList = append(*resultList,OneList)
		}

		Sum = Sum+root.Val
		list = list[:len(list)-1]
	}else if (root.Left.Val<Sum){

	}


}
