package graph

// Graph defines the core graph interface
type Graph interface {
	AddNode(id int) error                       // 添加節點
	AddEdge(from, to int, weight float64) error // 添加邊
	GetNeighbors(node int) ([]Edge, error)      // 獲取某節點的鄰居節點列表
	IsDirected() bool                           // 判斷圖是否為有向圖
	IsWeighted() bool                           // 判斷圖是否為加權圖
	NodeCount() int                             // 返回圖中節點的數量
	EdgeCount() int                             // 返回圖中邊的數量
	GetNodes() []int                            // 返回圖中所有節點
	GetEdges(node int) ([]Edge, error)          // 返回指定節點的邊列表
}

// Edge represents a graph edge
type Edge struct {
	To     int     // 終點節點
	Weight float64 // 邊的權重
}
