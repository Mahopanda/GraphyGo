package graph

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
