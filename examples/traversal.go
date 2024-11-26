package main

import (
	"fmt"

	"github.com/Mahopanda/GraphyGo/pkg/graph"
)

// 遍历算法
func main() {
	// 创建一个有向图
	g := graph.NewAdjacencyList(true, false)

	// 添加节点和边
	g.AddNode(1)
	g.AddNode(2)
	g.AddNode(3)
	g.AddEdge(1, 2, 0)
	g.AddEdge(2, 3, 0)
	g.AddEdge(1, 3, 0)

	// 广度优先遍历
	bfsResult, _ := graph.BFS(g, 1)
	fmt.Println("广度优先遍历结果:", bfsResult) // 输出: [1, 2, 3]

	// 深度优先遍历
	dfsResult, _ := graph.DFS(g, 1)
	fmt.Println("深度优先遍历结果:", dfsResult) // 输出: [1, 2, 3]
}
