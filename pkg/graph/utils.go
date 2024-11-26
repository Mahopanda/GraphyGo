package graph

import (
	"fmt"
	"strings"
)

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

// ToPlantUML generates a PlantUML-compatible string representation of the graph.
//
// Parameters:
// - g: The graph to represent (must implement the Graph interface).
//
// Returns:
// - A string in PlantUML format that can be rendered to visualize the graph.
//
// Example:
// g := NewAdjacencyList(true, false)
// g.AddNode(1)
// g.AddNode(2)
// g.AddEdge(1, 2, 0)
// uml := ToPlantUML(g)
// fmt.Println(uml)
// @startuml
// 1 --> 2
// @enduml
func ToPlantUML(g Graph) string {
	var sb strings.Builder
	sb.WriteString("@startuml\n")

	// 添加有向或无向的边
	for from := range g.(*AdjacencyList).edges {
		edges, _ := g.GetNeighbors(from)
		for _, edge := range edges {
			// 使用方向性表示箭头
			if g.IsDirected() {
				sb.WriteString(fmt.Sprintf("%d --> %d\n", from, edge.To))
			} else {
				sb.WriteString(fmt.Sprintf("%d -- %d\n", from, edge.To))
			}
		}
	}

	sb.WriteString("@enduml\n")
	return sb.String()
}
