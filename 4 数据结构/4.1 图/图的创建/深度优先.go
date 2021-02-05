package main

import "fmt"

// 得到第一个邻接节点的下标w
func getFirstNeighbor(index int)int  {
	for j:=0;j<len(vertextList);j++ {
		if edges[index][j]>0{
			return j
		}
	}
	return -1
}
// 根据第一个邻接点的下标获取下一个临接点
func getNextNeighbor(v1 int,v2 int) int {
	for j:=v2+1;j<len(vertextList);j++ {
		if edges[v1][j]>0 {
			return j
		}
	}
	return -1
}

func dfs(isVisited []bool,i int)  {
	// 首先我们访问该节点输出
	fmt.Print(getValueByid(i))
	// 把节点设置为已经访问
	isVisited[i] = true
	// 查找节点i的第一个临接节点w
	w := getFirstNeighbor(i)
	for w!=-1 { //说明有
		if !isVisited[w]{
			dfs(isVisited,w)
		}
		// 如果w节点已经被访问过
		w = getNextNeighbor(i,w)
	}
}

func UseDfs()  {
	isVisited = make([]bool,len(vertextList))
	// 遍历所有的节点，进行dfs回溯
	for i:=0;i<getNumberOfVertex();i++ {
		if !isVisited[i]{
			dfs(isVisited,i)
		}
	}
}