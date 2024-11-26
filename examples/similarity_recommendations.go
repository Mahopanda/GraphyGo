package main

import (
	"fmt"
	"pkg/graph"
)

//基于标签的相似度推荐

func main() {
	// 创建商品图
	pg := graph.NewProductGraph(true, true)

	// 添加商品
	products := []*graph.Product{
		{ID: 1, Labels: []string{"electronics", "phone"}},
		{ID: 2, Labels: []string{"electronics", "laptop"}},
		{ID: 3, Labels: []string{"electronics", "tablet"}},
		{ID: 4, Labels: []string{"furniture", "chair"}},
		{ID: 5, Labels: []string{"furniture", "desk"}},
	}
	for _, product := range products {
		pg.AddProduct(product)
	}

	// 添加相似度
	pg.AddSimilarity(1, 2, 0.8) // 手机和笔记本的相似度
	pg.AddSimilarity(1, 3, 0.6) // 手机和平板的相似度
	pg.AddSimilarity(2, 3, 0.9) // 笔记本和平板的相似度
	pg.AddSimilarity(4, 5, 0.7) // 椅子和桌子的相似度

	// 获取推荐
	recommendations, err := pg.GetRecommendations(1, 2) // 获取与商品1相关的前2个商品
	if err != nil {
		fmt.Printf("获取推荐失败: %v\n", err)
		return
	}

	// 打印推荐结果
	fmt.Println("推荐商品:")
	for _, rec := range recommendations {
		fmt.Printf("商品ID: %d, 标签: %v\n", rec.ID, rec.Labels)
	}
}
