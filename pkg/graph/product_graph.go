package graph

import "sort"

// 商品信息
type Product struct {
	ID     int      // 商品唯一标识
	Labels []string // 商品的标签
}

// ProductGraph 扩展图的结构以支持商品信息
type ProductGraph struct {
	Graph      *AdjacencyList   // 图结构
	ProductMap map[int]*Product // 商品信息映射
}

// NewProductGraph 创建一个商品图
func NewProductGraph(directed, weighted bool) *ProductGraph {
	return &ProductGraph{
		Graph:      NewAdjacencyList(directed, weighted),
		ProductMap: make(map[int]*Product),
	}
}

// AddProduct 添加商品节点
func (pg *ProductGraph) AddProduct(product *Product) error {
	err := pg.Graph.AddNode(product.ID)
	if err != nil {
		return err
	}
	pg.ProductMap[product.ID] = product
	return nil
}

// AddSimilarity 添加商品之间的相似度（边）
func (pg *ProductGraph) AddSimilarity(from, to int, similarity float64) error {
	return pg.Graph.AddEdge(from, to, similarity)
}

// GetRecommendations 获取推荐商品
func (pg *ProductGraph) GetRecommendations(productID int, topN int) ([]*Product, error) {
	neighbors, err := pg.Graph.GetNeighbors(productID)
	if err != nil {
		return nil, err
	}

	// 按权重排序获取前N个邻居
	sort.Slice(neighbors, func(i, j int) bool {
		return neighbors[i].Weight > neighbors[j].Weight
	})

	recommendations := []*Product{}
	for i, edge := range neighbors {
		if i >= topN {
			break
		}
		recommendations = append(recommendations, pg.ProductMap[edge.To])
	}

	return recommendations, nil
}
