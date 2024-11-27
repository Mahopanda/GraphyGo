package graph

import (
	"fmt"
	"math/rand"
)

// BFS performs a breadth-first traversal of the graph starting from the given node.
//
// Parameters:
// - g: The graph to traverse (must implement the Graph interface).
// - start: The starting node for the traversal.
//
// Returns:
// - A slice of integers representing the order of visited nodes.
// - An error if the traversal encounters issues (e.g., invalid start node).
//
// Example:
// g := NewAdjacencyList(true, false)
// g.AddNode(1)
// g.AddNode(2)
// g.AddEdge(1, 2, 0)
// result, err := BFS(g, 1)
// fmt.Println(result) // Output: [1, 2]

// 初始化訪問標記
// 將起點加入佇列
// 當佇列不為空：
//     取出佇列的第一個節點
//     如果該節點未訪問：
//         標記為已訪問
//         將其鄰居加入佇列
// 返回訪問過的節點列表

func BFS(g Graph, start int) ([]int, error) {
	visited := make(map[int]bool)
	queue := []int{start}
	result := []int{}

	for len(queue) > 0 { // 當佇列不為空時，繼續遍歷
		node := queue[0]
		queue = queue[1:]

		if visited[node] { // 如果節點已經訪問過，則跳過
			continue
		}
		visited[node] = true          // 標記節點為已訪問
		result = append(result, node) // 將節點添加到結果列表中

		neighbors, err := g.GetNeighbors(node)
		if err != nil {
			return nil, err // 獲取鄰居節點列表失敗，返回錯誤
		}
		for _, neighbor := range neighbors {
			if !visited[neighbor.To] {
				queue = append(queue, neighbor.To) // 將未訪問的鄰居節點添加到佇列中
			}
		}
	}
	return result, nil // 返回遍歷結果
}

// DFS performs a depth-first traversal of the graph starting from the given node.
func DFS(g Graph, start int) ([]int, error) {
	visited := make(map[int]bool)
	result := []int{}

	var dfs func(node int) error
	dfs = func(node int) error {
		visited[node] = true
		result = append(result, node)

		neighbors, err := g.GetNeighbors(node)
		if err != nil {
			return err
		}

		for _, neighbor := range neighbors {
			if !visited[neighbor.To] {
				if err := dfs(neighbor.To); err != nil {
					return err
				}
			}
		}
		return nil
	}

	if err := dfs(start); err != nil {
		return nil, err
	}
	return result, nil
}

// RandomWalk performs a random walk on the graph for a specified number of steps.
func RandomWalk(g *AdjacencyList, start int, steps int) ([]int, error) {
	// 驗證起始節點是否存在
	if !g.HasNode(start) {
		return nil, fmt.Errorf("start node %d does not exist", start)
	}

	// 初始化結果切片，包含起始節點
	walk := []int{start}
	current := start

	// 執行 steps-1 次移動（因為起始節點已經計入）
	for i := 0; i < steps-1; i++ {
		// 獲取當前節點的所有鄰居
		edges, err := g.GetNeighbors(current)
		if err != nil {
			return nil, fmt.Errorf("failed to get neighbors for node %d: %v", current, err)
		}

		// 轉換鄰居為節點列表
		neighbors := []int{}
		for _, edge := range edges {
			neighbors = append(neighbors, edge.To)
		}

		// 檢查是否有鄰居節點
		if len(neighbors) == 0 {
			return nil, fmt.Errorf("node %d has no neighbors", current)
		}

		// 隨機選擇下一個節點
		next := neighbors[rand.Intn(len(neighbors))]
		walk = append(walk, next)
		current = next
	}

	return walk, nil
}
