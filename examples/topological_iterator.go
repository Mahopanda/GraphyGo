package main

import (
	"fmt"

	"github.com/Mahopanda/GraphyGo/pkg/graph"
)

// 拓扑排序用于有向无环图（DAG），按依赖关系排序节点。例如，课程安排中，某些课程必须先完成才能上后续课程。
// TopologicalIterator 使用拓扑排序遍历图
func TopologicalIterator(g *graph.AdjacencyList) ([]int, error) {
	inDegree := make(map[int]int)
	order := []int{}

	// 計算每個節點的入度
	for _, from := range g.GetNodes() {
		edges, err := g.GetEdges(from)
		if err != nil {
			return nil, fmt.Errorf("failed to get edges for node %d: %v", from, err)
		}
		for _, edge := range edges {
			inDegree[edge.To]++
		}
	}

	// 初始化入度為 0 的節點隊列
	queue := []int{}
	for _, node := range g.GetNodes() {
		if inDegree[node] == 0 {
			queue = append(queue, node)
		}
	}

	// BFS 模擬拓撲排序
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		order = append(order, current)

		// 減少鄰居的入度
		edges, err := g.GetEdges(current)
		if err != nil {
			return nil, fmt.Errorf("failed to get edges for node %d: %v", current, err)
		}
		for _, edge := range edges {
			inDegree[edge.To]--
			if inDegree[edge.To] == 0 {
				queue = append(queue, edge.To)
			}
		}
	}

	// 如果排序結果數量小於節點數，說明圖中有環
	if len(order) != len(g.GetNodes()) {
		return nil, fmt.Errorf("graph contains a cycle, topological sort is not possible")
	}

	return order, nil
}

// DetectCycle 使用DFS檢測圖中是否存在環
func DetectCycle(g *graph.AdjacencyList) bool {
	visited := make(map[int]bool)
	recStack := make(map[int]bool)

	var dfs func(node int) bool
	dfs = func(node int) bool {
		if recStack[node] { // 當前節點已經在遞歸堆疊中，表示存在環
			return true
		}
		if visited[node] { // 節點已被訪問過，無需再次檢查
			return false
		}

		visited[node] = true
		recStack[node] = true

		// 遍歷鄰居節點
		edges, _ := g.GetEdges(node)
		for _, edge := range edges {
			if dfs(edge.To) {
				return true
			}
		}

		recStack[node] = false
		return false
	}

	for _, node := range g.GetNodes() {
		if dfs(node) {
			return true
		}
	}

	return false
}

func PrintGraph(g *graph.AdjacencyList) {
	fmt.Println("Graph structure:")
	for _, node := range g.GetNodes() {
		edges, _ := g.GetEdges(node)
		fmt.Printf("Node %d ->", node)
		for _, edge := range edges {
			fmt.Printf(" %d", edge.To)
		}
		fmt.Println()
	}
}

func main() {
	// 創建一個有向圖（DAG）
	g := graph.NewAdjacencyList(true, false)

	// 添加節點（課程編號）
	courses := []int{1, 2, 3, 4, 5, 6}
	for _, course := range courses {
		g.AddNode(course)
	}

	// 添加邊（課程依賴關係）
	prerequisites := []struct {
		from, to int
	}{
		{1, 3}, // 課程 1 是課程 3 的前置
		{2, 3}, // 課程 2 是課程 3 的前置
		{3, 4}, // 課程 3 是課程 4 的前置
		{3, 5}, // 課程 3 是課程 5 的前置
		{4, 6}, // 課程 4 是課程 6 的前置
		// {6, 3}, // 刪除此邊，避免環
	}
	for _, edge := range prerequisites {
		g.AddEdge(edge.from, edge.to, 0)
	}

	// 打印圖的結構
	PrintGraph(g)

	// 環檢測
	if DetectCycle(g) {
		fmt.Println("圖中存在環，無法進行拓撲排序")
		return
	}

	// 運行拓撲排序
	order, err := TopologicalIterator(g)
	if err != nil {
		fmt.Printf("拓撲排序失敗: %v\n", err)
		return
	}

	// 輸出排序結果
	fmt.Println("拓撲排序結果 (課程安排順序):")
	fmt.Println(order)
}
