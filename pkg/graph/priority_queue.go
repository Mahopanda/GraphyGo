package graph

import (
	"container/heap"
)

// Item 定義優先隊列中的元素
type Item struct {
	value    int     // 節點ID
	priority float64 // 優先級(距離)
	index    int     // 在堆中的索引
}

// PriorityQueue 實現了一個最小堆的優先隊列
type PriorityQueue []*Item

// Len 返回隊列長度
func (pq PriorityQueue) Len() int { return len(pq) }

// Less 比較優先級
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

// Swap 交換元素位置
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push 添加新元素
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

// Pop 移除並返回最小元素
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // 避免記憶體洩漏
	item.index = -1 // 標記為已移除
	*pq = old[0 : n-1]
	return item
}

// Update 更新元素的優先級
func (pq *PriorityQueue) Update(item *Item, value int, priority float64) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
} 