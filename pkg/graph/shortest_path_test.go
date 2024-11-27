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
