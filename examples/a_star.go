package main

import (
	"fmt"
	"math"

	"github.com/Mahopanda/GraphyGo/pkg/graph"
)

func main() {
	// 創建帶權重的有向圖
	g := graph.NewAdjacencyList(true, true)

	// 添加節點
	nodes := []int{1, 2, 3, 4, 5, 6}
	for _, node := range nodes {
		g.AddNode(node)
	}

	// 添加邊 (起點, 終點, 權重)
	edges := []struct {
		from, to int
		weight   float64
	}{
		{1, 2, 2},  // 1 → 2，距離 2
		{1, 3, 4},  // 1 → 3，距離 4
		{2, 4, 7},  // 2 → 4，距離 7
		{3, 4, 1},  // 3 → 4，距離 1
		{4, 5, 3},  // 4 → 5，距離 3
		{5, 6, 1},  // 5 → 6，距離 1
		{3, 6, 10}, // 3 → 6，距離 10
	}
	for _, edge := range edges {
		err := g.AddEdge(edge.from, edge.to, edge.weight)
		if err != nil {
			fmt.Printf("Error adding edge (%d -> %d): %v\n", edge.from, edge.to, err)
			return
		}
	}

	// 定義啟發式函數
	// 假設每個節點的估計代價為固定值
	heuristic := func(current, goal int) float64 {
		// 模擬簡單的啟發式，根據節點 ID 的差異
		// 更接近實際場景可以使用歐幾里得��離等更精確的啟發式函數
		return math.Abs(float64(goal - current))
	}

	// 運行 A* 算法
	startNode := 1
	goalNode := 6
	path, err := graph.AStar(g, startNode, goalNode, heuristic)
	if err != nil {
		fmt.Printf("Error running A*: %v\n", err)
		return
	}

	// 輸出結果
	fmt.Printf("從節點 %d 到節點 %d 的最短路徑: %v\n", startNode, goalNode, path)
}
