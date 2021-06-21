package main
// 二叉树节点间的最大距离问题
//限定语言：Python、C++、Java、Go、C、Javascript、Python 3
//从二叉树的节点 A 出发，可以向上或者向下走，但沿途的节点只能经过一次，当到达节点 B 时，路径上的节点数叫作 A 到 B 的距离。
//现在给出一棵二叉树，求整棵树上每对节点之间的最大距离。
//示例1
//输入
//{1,2,3,4,5,6,7}
//输出
//5


type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
}

func maxDistance(root *TreeNode) int {
	result := 0

	type Digui func(root *TreeNode) int
	var DiGuiFunc Digui
	DiGuiFunc = func(root *TreeNode) int {
		// 递归终止条件
		if root == nil {
		return 0
	}
		// 递归过程
		leftMaxDepth := DiGuiFunc(root.Left)
		rightMaxDepth := DiGuiFunc(root.Right)
		result = max(result,leftMaxDepth+rightMaxDepth+1)
		// 左右子树的最高深度+1
		return 1 + max(leftMaxDepth, rightMaxDepth)
	}

	DiGuiFunc(root)
	return result

}



func max(v1, v2 int) int {
	if v1 >= v2 {
		return v1
	}
	return v2
}
