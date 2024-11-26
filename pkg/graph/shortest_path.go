package graph

import (
	"container/heap"
	"fmt"
	"math"
)

// Dijkstra 實現Dijkstra最短路徑算法
func Dijkstra(g Graph, start int) (distances map[int]float64, predecessors map[int]int, err error) {
	if !g.IsWeighted() {
		return nil, nil, fmt.Errorf("Dijkstra requires a weighted graph")
	}

	distances = make(map[int]float64)
	predecessors = make(map[int]int)
	visited := make(map[int]bool)

	// 初始化距離為無窮大
	for node := 0; node < g.NodeCount(); node++ {
		distances[node] = math.Inf(1)
	}
	distances[start] = 0

	// 創建優先隊列
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Item{
		value:    start,
		priority: 0,
	})

	for pq.Len() > 0 {
		// 獲取當前最短距離的節點
		item := heap.Pop(&pq).(*Item)
		u := item.value

		if visited[u] {
			continue
		}
		visited[u] = true

		// 檢查所有相鄰節點
		neighbors, err := g.GetNeighbors(u)
		if err != nil {
			return nil, nil, err
		}

		for _, edge := range neighbors {
			v := edge.To
			if !visited[v] {
				alt := distances[u] + edge.Weight
				if alt < distances[v] {
					distances[v] = alt
					predecessors[v] = u
					heap.Push(&pq, &Item{
						value:    v,
						priority: alt,
					})
				}
			}
		}
	}

	return distances, predecessors, nil
}
