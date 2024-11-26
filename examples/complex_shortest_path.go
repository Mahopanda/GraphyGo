package main

import (
	"fmt"

	"github.com/Mahopanda/GraphyGo/pkg/graph"
)

// 背景：
// 该图表示一个城市中的道路网络，节点代表位置（如交叉路口），边代表道路。边的权重表示道路的通行代价（如时间、距离或拥堵程度）。部分道路是单向的，部分是双向的。

// 任务：
// 1. 构建城市的道路网络图。
// 2. 计算从指定起点到所有其他节点的最短路径。
// 3. 可视化图形并验证最短路径结果。

func main() {
	// 创建一个带权图
	g := graph.NewAdjacencyList(true, true)

	// 添加更多节点 (0-7)
	for i := 0; i < 8; i++ {
		g.AddNode(i)
	}

	// 添加各种类型的边
	// 1. 双向道路
	addBidirectionalEdge(g, 0, 1, 2) // 主干道
	addBidirectionalEdge(g, 1, 2, 3) // 主干道
	addBidirectionalEdge(g, 2, 3, 1) // 主干道

	// 2. 单向道路
	g.AddEdge(1, 4, 2) // 单向快速通道
	g.AddEdge(4, 3, 1) // 单向快速通道

	// 3. 环形路径
	addBidirectionalEdge(g, 3, 5, 2)
	addBidirectionalEdge(g, 5, 6, 3)
	addBidirectionalEdge(g, 6, 3, 4)

	// 4. 平行边（不同权重的alternative路径）
	g.AddEdge(0, 2, 7) // 较长的直接路径
	g.AddEdge(1, 3, 6) // 较长的直接路径

	// 5. 添加一个几乎孤立的节点（只有一条连接）
	g.AddEdge(6, 7, 1)

	// 测试从不同起点的最短路径
	testShortestPaths(g, 0) // 从起点0开始
	testShortestPaths(g, 3) // 从起点3开始

	// 输出图的可视化
	uml, err := graph.ToPlantUML(g)
	if err != nil {
		fmt.Printf("生成 PlantUML 失败: %v\n", err)
		return
	}
	fmt.Println("\nPlantUML 图形表示:")
	fmt.Println(uml)
}

// 添加双向边的辅助函数
func addBidirectionalEdge(g *graph.AdjacencyList, from, to int, weight float64) {
	g.AddEdge(from, to, weight)
	g.AddEdge(to, from, weight)
}

// 测试并打印最短路径
func testShortestPaths(g *graph.AdjacencyList, start int) {
	distances, predecessors, err := graph.Dijkstra(g, start)
	if err != nil {
		fmt.Printf("计算最短路径时发生错误: %v\n", err)
		return
	}

	fmt.Printf("\n从节点 %d 出发的最短路径结果:\n", start)
	fmt.Println("距离:", distances)
	fmt.Println("前驱节点:", predecessors)

	// 为每个节点重建并打印完整路径
	for end := 0; end < 8; end++ {
		if end != start {
			path := reconstructPath(predecessors, start, end)
			fmt.Printf("到节点 %d 的路径: %v (距离: %.1f)\n",
				end, path, distances[end])
		}
	}
}

// 重建完整路径
func reconstructPath(predecessors map[int]int, start, end int) []int {
	path := []int{end}
	for current := end; current != start; {
		prev, exists := predecessors[current]
		if !exists {
			return []int{} // 无法到达的路径
		}
		path = append([]int{prev}, path...)
		current = prev
	}
	return path
}
