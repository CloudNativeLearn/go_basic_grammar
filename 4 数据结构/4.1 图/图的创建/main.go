package main

import "fmt"

func main() {
	// 测试一把是否ok
	n := 5
	vertextList = []string{"A", "B", "C", "D", "E"}
	Graph(n)

	// 显示矩阵
	getAllValue()

	// 添加边
	insertEdge(0, 1, 1) // A-B
	insertEdge(0, 2, 1) // A-C
	insertEdge(1, 2, 1) // B-C
	insertEdge(1, 3, 1) // B-D
	insertEdge(1, 4, 1) // B-E

	fmt.Println("插入数据之后再显示")
	// 显示矩阵
	getAllValue()

}

var vertextList []string // 存储顶点集合
var edges [][]int        // 存储图对应的邻接矩阵
var numOfEdges int       // 表示边的数目
var isVisited[] bool       // 表示读取哪些已读节点

// 构造器
func Graph(n int) {
	// 初始化矩阵和VertextList
	//edges =  [][]int{}
	edges = make([][]int, n)
	// 初始化切片
	for i := range edges {
		edges[i] = make([]int, n)
	}

	vertextList = make([]string, n)
	isVisited = make([]bool,n)
	numOfEdges = 0

}

// 插入节点
func insertVertex(vertex string) {
	vertextList = append(vertextList, vertex)
}

// 添加边
func insertEdge(A int, B int, widge int) {
	edges[A][B] = widge
	edges[B][A] = widge
	numOfEdges = numOfEdges + 1
}

// 返回节点的个数
func getNumberOfVertex() int {
	return len(vertextList)
}

// 返回边的数目
func getNumOfEdges() int {
	return numOfEdges
}

// 根据id返回对应的value 
func getValueByid(i int) string {
	return vertextList[i]
}

// 返回v1 v2对应的权值
func getWeightByid(A int, B int) int {
	return edges[A][B]
}

// 显示所有
func getAllValue() {
	for _, value := range edges {
		//fmt.Println(k)
		fmt.Println(value)
		//fmt.Println("--------------")
	}
}
