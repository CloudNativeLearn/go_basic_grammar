package main

// 节点间通路。给定有向图，设计一个算法，找出两个节点之间是否存在一条路径。
//
//示例1:
//
// 输入：n = 3, graph = [[0, 1], [0, 2], [1, 2], [1, 2]], start = 0, target = 2
// 输出：true
//示例2:
//
// 输入：n = 5, graph = [[0, 1], [0, 2], [0, 4], [0, 4], [0, 1], [1, 3], [1, 4], [1, 3], [2, 3], [3, 4]], start = 0, target = 4
// 输出 true
//提示：
//
//节点数量n在[0, 1e5]范围内。
//节点编号大于等于 0 小于 n。
//图中可能存在自环和平行边。
//通过次数12,151提交次数22,940
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/route-between-nodes-lcci
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。




func findWhetherExistsPath(n int, graph [][]int, start int, target int) bool {

	WhetherHasGo := make([]bool, n)



	// 判断在一个数组中 是否有某一对数字存在
	FindInarr := func(start int,end int) bool {
		for _,v := range graph{
			if v[0]==start && v[1]==end{
				return true
			}
		}
		return false
	}


	// 创建函数 进行深度遍历
	type DfsFunc func(int) bool
	var Dfs DfsFunc
	Dfs = func(ant int) bool {
		//循环所有的节点
		for i := 0; i < n; i++ {
			if WhetherHasGo[i] == false && FindInarr(ant,i) {
				// 没有去过而且可达
				if i == target {
					return true
				}
				WhetherHasGo[i] = true
				if Dfs(i){
					return true
				}
			}
		}




		return false
	}
	return Dfs(start)

}
