package main

import (
	"container/heap"
	"fmt"
	"math"

	"github.com/Mahopanda/GraphyGo/pkg/graph"
)

// Dijkstra 使用優先隊列實現單源最短路徑算法
func Dijkstra(g *graph.AdjacencyList, start int) (map[int]float64, map[int]int, error) {
	// 確保起始節點存在於圖中
	if !g.HasNode(start) {
		return nil, nil, fmt.Errorf("start node %d does not exist in the graph", start)
	}

	// 初始化距離和前驅節點
	distances := make(map[int]float64) // 節點到起點的最短距離
	predecessors := make(map[int]int)  // 前驅節點
	for _, node := range g.GetNodes() {
		distances[node] = math.Inf(1) // 初始距離為無窮大
		predecessors[node] = -1       // 前驅節點初始化為 -1
	}
	distances[start] = 0 // 起點到自己的距離為 0

	// 初始化優先隊列
	pq := graph.NewPriorityQueue()
	heap.Push(pq, graph.NewItem(start, 0))

	// 開始計算最短路徑
	for pq.Len() > 0 {
		// 取出當前距離最小的節點
		current := graph.GetItemValue(heap.Pop(pq).(*graph.Item))

		// 遍歷當前節點的所有鄰居
		edges, err := g.GetEdges(current)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to get edges for node %d: %v", current, err)
		}

		for _, edge := range edges {
			alt := distances[current] + edge.Weight // 計算新距離
			if alt < distances[edge.To] {
				// 更新距離和前驅節點
				distances[edge.To] = alt
				predecessors[edge.To] = current

				// 更新優先隊列
				heap.Push(pq, graph.NewItem(edge.To, alt))
			}
		}
	}

	return distances, predecessors, nil
}

func main() {
	// 創建帶權重的有向圖
	g := graph.NewAdjacencyList(true, true)

	// 添加節點
	nodes := []int{1, 2, 3, 4, 5}
	for _, node := range nodes {
		g.AddNode(node)
	}

	// 添加邊 (起點, 終點, 權重)
	edges := []struct {
		from, to int
		weight   float64
	}{
		{1, 2, 1}, // 節點 1 到 2，權重為 1
		{1, 3, 4}, // 節點 1 到 3，權重為 4
		{2, 3, 2}, // 節點 2 到 3，權重為 2
		{2, 4, 7}, // 節點 2 到 4，權重為 7
		{3, 5, 3}, // 節點 3 到 5，權重為 3
		{4, 5, 1}, // 節點 4 到 5，權重為 1
	}
	for _, edge := range edges {
		err := g.AddEdge(edge.from, edge.to, edge.weight)
		if err != nil {
			fmt.Printf("Error adding edge (%d -> %d): %v\n", edge.from, edge.to, err)
			return
		}
	}

	// 運行 Dijkstra 算法
	startNode := 1
	distances, predecessors, err := Dijkstra(g, startNode)
	if err != nil {
		fmt.Printf("Error running Dijkstra: %v\n", err)
		return
	}

	// 輸出最短路徑結果
	fmt.Println("最短路徑距離:")
	for node, dist := range distances {
		fmt.Printf("Node %d: %.1f\n", node, dist)
	}

	fmt.Println("前驅節點:")
	for node, pred := range predecessors {
		fmt.Printf("Node %d: %d\n", node, pred)
	}
}
