package graph

import (
	"errors"
	"fmt"
	"math"

	"container/heap"
)

type AdjacencyList struct {
	directed bool           // 是否為有向圖
	weighted bool           // 是否為加權圖
	nodes    map[int]bool   // 節點列表
	edges    map[int][]Edge // 邊列表
}

// NewAdjacencyList creates a new graph using an adjacency list representation.
//
// Parameters:
// - directed: If true, the graph is directed. Otherwise, it's undirected.
// - weighted: If true, the graph supports edge weights.
//
// Returns:
// - An instance of AdjacencyList initialized with the given properties.
//
// Example:
// g := NewAdjacencyList(true, false)
// g.AddNode(1)
// g.AddEdge(1, 2, 0)

func NewAdjacencyList(directed, weighted bool) *AdjacencyList {
	return &AdjacencyList{
		directed: directed,
		weighted: weighted,
		nodes:    make(map[int]bool),
		edges:    make(map[int][]Edge),
	}
}

func (g *AdjacencyList) AddNode(node int) error {
	if _, exists := g.nodes[node]; exists {
		return fmt.Errorf("節點 %d 已存在", node)
	}
	g.nodes[node] = true
	return nil
}

func (g *AdjacencyList) AddEdge(from, to int, weight float64) error {
	// 新增一條邊到圖中，若為無向圖會自動添加反向邊
	// 參數:
	// - from: 起始節點
	// - to: 終止節點
	// - weight: 邊的權重 (若為無權圖，權重必須為 0)
	// 回傳:
	// - error: 若節點不存在或在無權圖中設置了權重，則返回錯誤

	if !g.weighted && weight != 0 {
		return errors.New("weight not allowed in unweighted graph") // 無權圖中不允許設置權重
	}
	if _, exists := g.nodes[from]; !exists {
		return errors.New("from node does not exist") // 起始節點不存在
	}
	if _, exists := g.nodes[to]; !exists {
		return errors.New("to node does not exist") // 終止節點不存在
	}
	// 將邊添加到起始節點的鄰居列表中
	g.edges[from] = append(g.edges[from], Edge{To: to, Weight: weight}) // 將邊添加到起始節點的鄰居列表中
	if !g.directed {
		// 若為無向圖，添加反向邊
		g.edges[to] = append(g.edges[to], Edge{To: from, Weight: weight}) // 若為無向圖，添加反向邊
	}
	return nil
}

func (g *AdjacencyList) GetNeighbors(node int) ([]Edge, error) {
	// 獲取某節點的鄰居節點列表
	// 參數:
	// - node: 要查詢的節點
	// 回傳:
	// - []Edge: 鄰居節點的邊列表
	// - error: 若節點不存在，返回錯誤

	if _, exists := g.nodes[node]; !exists {
		return nil, fmt.Errorf("node %d does not exist in the graph", node) // 節點不存在
	}
	return g.edges[node], nil // 返回節點的鄰居節點列表
}

func (g *AdjacencyList) IsDirected() bool {
	// 判斷圖是否為有向圖
	// 回傳:
	// - true: 圖是有向圖
	// - false: 圖是無向圖
	return g.directed
}

func (g *AdjacencyList) IsWeighted() bool {
	// 判斷圖是否為加權圖
	// 回傳:
	// - true: 圖是加權圖
	// - false: 圖是非加權圖
	return g.weighted
}

func (g *AdjacencyList) NodeCount() int {
	// 返回圖中節點的數量
	return len(g.nodes)
}

func (g *AdjacencyList) EdgeCount() int {
	// 返回圖中邊的數量
	count := 0
	for _, neighbors := range g.edges {
		count += len(neighbors) // 計算所有邊的數量
	}
	if !g.directed {
		count /= 2 // 無向圖中每條邊會被計算兩次，因此需要除以2
	}
	return count
}

func (g *AdjacencyList) HasNode(id int) bool {
	// 判斷圖中是否存在某節點
	// 參數:
	// - id: 要查詢的節點
	// 回傳:
	// - true: 節點存在
	// - false: 節點不存在
	_, exists := g.nodes[id]
	return exists
}

func (g *AdjacencyList) HasEdge(from, to int) bool {
	// 判斷圖中是否存在某條邊
	// 參數:
	// - from: 起始節點
	// - to: 終止節點
	// 回傳:
	// - true: 邊存在
	// - false: 邊不存���
	for _, edge := range g.edges[from] {
		if edge.To == to {
			return true
		}
	}
	return false
}

func (g *AdjacencyList) RemoveNode(id int) error {
	// 從圖中移除某節點
	// 參數:
	// - id: 要移除的節點
	// 回傳:
	// - error: 若節點不存在，返回錯誤
	if _, exists := g.nodes[id]; !exists {
		return fmt.Errorf("node %d does not exist", id) // 節點不存在
	}
	delete(g.nodes, id) // 從節點列表中移除節點
	delete(g.edges, id) // 從邊列表中移除節點

	// 遍歷所有節點，從鄰居列表中移除與該節點相關的邊
	for node := range g.edges {
		neighbors := g.edges[node]
		for i := 0; i < len(neighbors); i++ {
			if neighbors[i].To == id {
				g.edges[node] = append(neighbors[:i], neighbors[i+1:]...) // 從鄰居列表中移除邊
				break
			}
		}
	}
	return nil
}

func (g *AdjacencyList) GetNodes() []int {
	nodes := make([]int, 0, len(g.nodes))
	for node := range g.nodes {
		nodes = append(nodes, node)
	}
	return nodes
}

// GetEdges 返回指定節點的邊列表
func (g *AdjacencyList) GetEdges(node int) ([]Edge, error) {
	if _, exists := g.nodes[node]; !exists {
		return nil, fmt.Errorf("node %d does not exist", node)
	}
	return g.edges[node], nil
}

func AStar(g *AdjacencyList, start, goal int, heuristic func(int, int) float64) ([]int, error) {
	// 初始化距離和前驅節點
	distances := make(map[int]float64)
	predecessors := make(map[int]int)
	for _, node := range g.GetNodes() {
		distances[node] = math.Inf(1)
		predecessors[node] = -1
	}
	distances[start] = 0

	// 初始化優先隊列
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{value: start, priority: heuristic(start, goal)})

	// 開始搜索
	for pq.Len() > 0 {
		current := heap.Pop(pq).(*Item).value

		// 如果已到達目標，停止搜索
		if current == goal {
			break
		}

		// 遍歷當前節點的鄰居
		edges, _ := g.GetEdges(current)
		for _, edge := range edges {
			alt := distances[current] + edge.Weight
			if alt < distances[edge.To] {
				// 更新距離和前驅節點
				distances[edge.To] = alt
				predecessors[edge.To] = current

				// 計算總優先級 = 實際距離 + 啟發式距離
				priority := alt + heuristic(edge.To, goal)
				heap.Push(pq, &Item{value: edge.To, priority: priority})
			}
		}
	}

	// 在搜索完成後，重建路徑
	path := []int{}
	current := goal
	for current != -1 {
		path = append([]int{current}, path...)
		current = predecessors[current]
	}

	// 檢查是否找到路徑
	if len(path) == 0 || path[0] != start {
		return nil, fmt.Errorf("無法找到從節點 %d 到節點 %d 的路徑", start, goal)
	}

	return path, nil
}
