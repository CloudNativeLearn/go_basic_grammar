package main

// 检查子树。你有两棵非常大的二叉树：T1，有几万个节点；T2，有几万个节点。设计一个算法，判断 T2 是否为 T1 的子树。
//
//如果 T1 有这么一个节点 n，其子树与 T2 一模一样，则 T2 为 T1 的子树，也就是说，从节点 n 处把树砍断，得到的树与 T2 完全相同。
//
//注意：此题相对书上原题略有改动。
//
//示例1:
//
// 输入：t1 = [1, 2, 3], t2 = [2]
// 输出：true
//示例2:
//
// 输入：t1 = [1, null, 2, 4], t2 = [3, 2]
// 输出：false
//提示：
//
//树的节点数目范围为[0, 20000]。
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/check-subtree-lcci
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func checkSubTree(t1 *TreeNode, t2 *TreeNode) bool {

	type isChildTree func(n1 *TreeNode, n2 *TreeNode) bool
	var isTree isChildTree
	isTree = func(n1 *TreeNode, n2 *TreeNode) bool {
		if n1 == nil && n2 == nil {
			return true
		}
		if n1 == nil || n2 == nil {
			return false
		}
		return n1.Val == n2.Val && isTree(n1.Left, n2.Left) && isTree(n1.Right, n2.Right)
	}

	type Dfs func(node *TreeNode) bool
	var dfs Dfs
	dfs = func(node *TreeNode) bool {
		if t2 == nil {
			return true
		}
		if node == nil {
			return false
		}
		return isTree(node, t2) || dfs(node.Left) || dfs(node.Right)
	}

	return dfs(t1)

}
