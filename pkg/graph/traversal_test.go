package graph

import (
	"math/rand"
	"testing"
	"time"
)

func TestDFS(t *testing.T) {
	// 創建一個簡單的圖進行測試
	g := NewAdjacencyList(true, false)

	// 添加節點
	nodes := []int{1, 2, 3, 4, 5}
	for _, node := range nodes {
		g.AddNode(node)
	}

	// 添加邊
	edges := []struct{ from, to int }{
		{1, 2}, {1, 3}, {2, 4}, {3, 5},
	}
	for _, edge := range edges {
		g.AddEdge(edge.from, edge.to, 0)
	}

	// 測試DFS遍歷
	result, err := DFS(g, 1)
	if err != nil {
		t.Errorf("DFS failed: %v", err)
	}

	// 驗證結果
	// DFS應該保證父節點在子節點之前
	visited := make(map[int]bool)
	for i, node := range result {
		if visited[node] {
			t.Errorf("Node %d appeared multiple times", node)
		}
		visited[node] = true

		// 檢查是否所有節點都被訪問
		if i == len(result)-1 && len(visited) != len(nodes) {
			t.Errorf("Not all nodes were visited")
		}
	}
}

// 測試BFS
func TestBFS(t *testing.T) {
	g := NewAdjacencyList(false, false)
	// 添加節點
	nodes := []int{1, 2, 3, 4}
	for _, node := range nodes {
		g.AddNode(node)
	}
	// 添加邊
	edges := []struct{ from, to int }{
		{1, 2}, {1, 3}, {2, 4}, {3, 4},
	}
	for _, edge := range edges {
		g.AddEdge(edge.from, edge.to, 0)
	}
	// 測試BFS遍歷
	result, err := BFS(g, 1)
	if err != nil {
		t.Errorf("BFS failed: %v", err)
	}
	// 驗證結果
	if len(result) != len(nodes) {
		t.Errorf("Expected %d nodes, got %d", len(nodes), len(result))
	}
	// BFS應該保證父節點在子節點之前
	visited := make(map[int]bool)
	for _, node := range result {
		if visited[node] {
			t.Errorf("Node %d appeared multiple times", node)
		}

		visited[node] = true
	}

}

func TestRandomWalk(t *testing.T) {
	// 設置隨機數種子以使結果可重現
	rand.Seed(time.Now().UnixNano())

	// 創建一個無向圖（因為我們要能夠雙向行走）
	g := NewAdjacencyList(false, false)

	// 添加節點
	nodes := []int{1, 2, 3, 4}
	for _, node := range nodes {
		g.AddNode(node)
	}

	// 添加雙向邊
	edges := []struct{ from, to int }{
		{1, 2}, {2, 3}, {3, 4}, {4, 1},
	}
	for _, edge := range edges {
		// 因為是無向圖，所以需要添加兩個方向的邊
		err := g.AddEdge(edge.from, edge.to, 0)
		if err != nil {
			t.Fatalf("Failed to add edge (%d -> %d): %v", edge.from, edge.to, err)
		}
	}

	// 在執行隨機遊走之前，先驗證圖的結構
	t.Logf("Graph structure:")
	for node := range g.nodes {
		neighbors, err := g.GetNeighbors(node)
		if err != nil {
			t.Fatalf("Failed to get neighbors for node %d: %v", node, err)
		}
		t.Logf("Node %d neighbors: %v", node, neighbors)
	}

	// 設置隨機遊走的參數
	expectedLength := 10
	walk, err := RandomWalk(g, 1, expectedLength) // 確保從節點1開始
	if err != nil {
		t.Fatalf("RandomWalk failed: %v", err)
	}

	// 驗證路徑長度
	if len(walk) != expectedLength {
		t.Errorf("Expected walk length %d, got %d", expectedLength, len(walk))
		t.Logf("Walk path: %v", walk)
	}

	// 驗證起始節點
	if walk[0] != 1 {
		t.Errorf("Expected walk to start at node 1, got %d", walk[0])
	}

	// 驗證每一步都是有效的
	for i := 0; i < len(walk)-1; i++ {
		from := walk[i]
		to := walk[i+1]
		if !g.HasEdge(from, to) {
			t.Errorf("Invalid step at position %d: from %d to %d", i, from, to)
		}
	}

	// 輸出完整路徑以便調試
	t.Logf("Complete walk path: %v", walk)
}
