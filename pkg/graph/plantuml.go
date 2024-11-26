package graph

import (
	"fmt"
	"strings"
)

// ToPlantUML generates a PlantUML-compatible string representation of the graph.
func ToPlantUML(g Graph) (string, error) {
	var sb strings.Builder
	sb.WriteString("@startuml\n")
	sb.WriteString("digraph G {\n")   // 使用 DOT 语言的有向图
	sb.WriteString("  rankdir=LR;\n") // 从左到右的布局

	// 添加节点
	for node := range g.(*AdjacencyList).nodes {
		sb.WriteString(fmt.Sprintf("  %d;\n", node))
	}

	// 添加边
	seenEdges := make(map[string]bool) // 防止重复添加无向边
	for from := range g.(*AdjacencyList).edges {
		edges, err := g.GetNeighbors(from)
		if err != nil {
			return "", err
		}
		for _, edge := range edges {
			edgeKey := fmt.Sprintf("%d-%d", from, edge.To)
			if !g.IsDirected() && seenEdges[edgeKey] {
				continue
			}
			sb.WriteString(fmt.Sprintf("  %d -> %d", from, edge.To))
			if g.IsWeighted() {
				sb.WriteString(fmt.Sprintf(" [label=\"%.1f\"]", edge.Weight))
			}
			sb.WriteString(";\n")
			if !g.IsDirected() {
				seenEdges[fmt.Sprintf("%d-%d", edge.To, from)] = true
			}
		}
	}

	sb.WriteString("}\n")
	sb.WriteString("@enduml\n")
	return sb.String(), nil
}
