package graph

import (
	"math"
	"testing"
)

func TestDijkstra(t *testing.T) {

	// 創建一個帶權重的有向圖
	g := NewAdjacencyList(true, true)

	// 添加節點
	for i := 0; i < 5; i++ {
		g.AddNode(i)
	}

	// 添加帶權重的邊
	edges := []struct {
		from   int
		to     int
		weight float64
	}{
		{0, 1, 4},
		{0, 2, 2},
		{1, 2, 1},
		{1, 3, 5},
		{2, 3, 8},
		{2, 4, 10},
		{3, 4, 2},
	}

	for _, edge := range edges {
		err := g.AddEdge(edge.from, edge.to, edge.weight)
		if err != nil {
			t.Fatalf("Failed to add edge: %v", err)
		}
	}

	// 測試從節點0開始的最短路徑
	distances, predecessors, err := Dijkstra(g, 0)
	if err != nil {
		t.Fatalf("Dijkstra failed: %v", err)
	}

	// 驗證最短距離
	expectedDistances := map[int]float64{
		0: 0,
		1: 4,
		2: 2,
		3: 9,
		4: 11,
	}

	for node, expectedDist := range expectedDistances {
		if math.Abs(distances[node]-expectedDist) > 1e-10 {
			t.Errorf("Incorrect distance for node %d: got %f, want %f",
				node, distances[node], expectedDist)
		}
	}

	// 驗證前驅節點
	expectedPredecessors := map[int]int{
		1: 0,
		2: 0,
		3: 1,
		4: 3,
	}

	for node, expectedPred := range expectedPredecessors {
		if predecessors[node] != expectedPred {
			t.Errorf("Incorrect predecessor for node %d: got %d, want %d",
				node, predecessors[node], expectedPred)
		}
	}
}

func TestDijkstraWithVisualization(t *testing.T) {
	// 创建一个带权图
	g := NewAdjacencyList(true, true)

	// 添加节点
	g.AddNode(0)
	g.AddNode(1)
	g.AddNode(2)
	g.AddNode(3)

	// 添加边
	g.AddEdge(0, 1, 1)
	g.AddEdge(1, 2, 2)
	g.AddEdge(0, 2, 4)
	g.AddEdge(2, 3, 1)

	// 生成可视化文件
	_, err := ToPlantUML(g)
	if err != nil {
		t.Fatalf("Failed to visualize graph: %v", err)
	}

	// 验证最短路径
	distances, _, err := Dijkstra(g, 0)

	if err != nil {
		t.Fatalf("Dijkstra failed: %v", err)
	}

	expectedDistances := map[int]float64{
		0: 0, // 到自身的距離為0
		1: 1, // 0->1 權重為1
		2: 3, // 0->1->2 權重為1+2=3
		3: 4, // 0->1->2->3 權重為1+2+1=4
	}
	for node, expectedDist := range expectedDistances {
		if math.Abs(distances[node]-expectedDist) > 1e-10 {
			t.Errorf("Incorrect distance for node %d: got %f, want %f",
				node, distances[node], expectedDist)
		}
	}
}


// 測試Dijkstra - 迷宮中的最短路徑
func TestDijkstra_MazeShortestPath(t *testing.T) {
	// 創建一個圖，表示迷宮中的節點和路徑
	g := NewAdjacencyList(true, true)

	// 添加節點（迷宮中的房間）
	nodes := []int{1, 2, 3, 4}
	for _, node := range nodes {
		g.AddNode(node)
	}

	// 添加邊（迷宮中的通路）和權重（表示距離）
	edges := []struct {
		from, to int
		weight   float64
	}{
		{1, 2, 1}, // 房間 1 → 房間 2，距離為 1
		{1, 3, 2}, // 房間 1 → 房間 3，距離為 2
		{2, 4, 1}, // 房間 2 → 房間 4，距離為 1
		{3, 4, 3}, // 房間 3 → 房間 4，距離為 3
	}

	// 添加邊並驗證
	for _, edge := range edges {
		err := g.AddEdge(edge.from, edge.to, edge.weight)
		if err != nil {
			t.Fatalf("Failed to add edge from %d to %d: %v", 
				edge.from, edge.to, err)
		}
	}

	// 驗證圖的結構
	t.Logf("Graph structure before Dijkstra:")
	for node := range g.nodes {
		edges := g.edges[node]
		t.Logf("Node %d edges: %v", node, edges)
	}

	// 使用Dijkstra算法查找從房間 1 到其他房間的最短路徑
	distances, predecessors, err := Dijkstra(g, 1)
	if err != nil {
		t.Fatalf("Dijkstra failed: %v", err)
	}

	// 輸出實際得到的結果以便調試
	t.Logf("Actual distances: %v", distances)
	t.Logf("Actual predecessors: %v", predecessors)

	// 驗證最短路徑的距離
	expectedDistances := map[int]float64{
		1: 0, // 到自己的距離為 0
		2: 1, // 房間 1 → 房間 2，距離 1
		3: 2, // 房間 1 → 房間 3，距離 2
		4: 2, // 房間 1 → 房間 2 → 房間 4，總距離 2
	}

	for node, expectedDist := range expectedDistances {
		actualDist := distances[node]
		if math.Abs(actualDist - expectedDist) > 1e-10 {
			t.Errorf("Distance mismatch for node %d: expected %.1f, got %.1f", 
				node, expectedDist, actualDist)
		}
	}

	// 驗證最短路徑的前驅節點
	expectedPredecessors := map[int]int{
		2: 1, // 房間 2 的前驅節點是房間 1
		3: 1, // 房間 3 的前驅節點是房間 1
		4: 2, // 房間 4 的前驅節點是房間 2
	}
	for node, expectedPred := range expectedPredecessors {
		if predecessors[node] != expectedPred {
			t.Errorf("Expected predecessor of node %d: %d, got %d", node, expectedPred, predecessors[node])
		}
	}

	// 驗證路徑重建功能（從前驅節點重建完整路徑）
	start, target := 1, 4
	expectedPath := []int{1, 2, 4}
	path := reconstructPath(predecessors, start, target)
	if len(path) != len(expectedPath) {
		t.Errorf("Expected path length %d, got %d", len(expectedPath), len(path))
	}
	for i, node := range path {
		if node != expectedPath[i] {
			t.Errorf("Expected node at position %d: %d, got %d", i, expectedPath[i], node)
		}
	}
}

// 重建路徑的輔助函數
func reconstructPath(predecessors map[int]int, start, target int) []int {
	path := []int{target}
	for current := target; current != start; {
		prev, exists := predecessors[current]
		if !exists {
			return []int{} // 無法到達的節點
		}
		path = append([]int{prev}, path...)
		current = prev
	}
	return path
}
