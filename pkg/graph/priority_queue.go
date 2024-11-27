package graph

import (
	"container/heap"
)

// Item 定義優先隊列中的元素
type Item struct {
	value    int     // 節點值
	priority float64 // 優先級（在Dijkstra算法中表示距離）
	index    int     // 在堆中的索引
}

// PriorityQueue 實現了heap.Interface
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

// Update 更新元素的優先級
func (pq *PriorityQueue) Update(item *Item, value int, priority float64) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

// NewItem 創建一個新的優先隊列項目
func NewItem(value int, priority float64) *Item {
	return &Item{
		value:    value,
		priority: priority,
		index:    -1, // 初始化為-1，表示尚未加入隊列
	}
}

// NewPriorityQueue 創建一個新的優先隊列
func NewPriorityQueue() *PriorityQueue {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	return &pq
}

// GetItemValue 獲取項目的值
func GetItemValue(item *Item) int {
	return item.value
}
