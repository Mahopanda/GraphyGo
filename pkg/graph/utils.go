package graph

// IsDAG checks whether the given graph is a Directed Acyclic Graph (DAG).
//
// Parameters:
// - g: The graph to check (must implement the Graph interface).
//
// Returns:
// - true if the graph is a DAG.
// - false if the graph contains a cycle.
//
// Example:
// g := NewAdjacencyList(true, false)
// g.AddNode(1)
// g.AddNode(2)
// g.AddEdge(1, 2, 0)
// isDAG := IsDAG(g)
// fmt.Println(isDAG) // Output: true
func IsDAG(g Graph) bool {
	visited := make(map[int]bool)
	recursionStack := make(map[int]bool)

	var dfs func(node int) bool
	dfs = func(node int) bool {
		visited[node] = true
		recursionStack[node] = true

		neighbors, _ := g.GetNeighbors(node)
		for _, edge := range neighbors {
			if !visited[edge.To] && dfs(edge.To) {
				return true
			} else if recursionStack[edge.To] {
				return true
			}
		}

		recursionStack[node] = false
		return false
	}

	for node := range visited {
		if !visited[node] && dfs(node) {
			return false
		}
	}
	return true
}
