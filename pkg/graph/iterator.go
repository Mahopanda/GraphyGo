package graph

// Iterator 定義圖遍歷的基本接口
type Iterator interface {
	// HasNext 返回是否還有下一個節點
	HasNext() bool
	// Next 返回下一個節點
	Next() (int, error)
	// Reset 重置迭代器
	Reset()
}

// DFSIterator 深度優先遍歷迭代器
type DFSIterator struct {
	graph     Graph
	start     int
	stack     []int
	visited   map[int]bool
	finished  bool
}

// BFSIterator 廣度優先遍歷迭代器
type BFSIterator struct {
	graph     Graph
	start     int
	queue     []int
	visited   map[int]bool
	finished  bool
}

// TopologicalIterator 拓撲排序迭代器
type TopologicalIterator struct {
	graph           Graph
	sorted          []int
	currentIndex    int
	visited         map[int]bool
	temporaryMark   map[int]bool
}

// RandomWalkIterator 隨機遊走迭代器
type RandomWalkIterator struct {
	graph         Graph
	current       int
	steps        int
	maxSteps     int
}

// ClosestFirstIterator 最近優先迭代器(用於Dijkstra)
type ClosestFirstIterator struct {
	graph         Graph
	start         int
	distances     map[int]float64
	visited       map[int]bool
	pq            *PriorityQueue
} 