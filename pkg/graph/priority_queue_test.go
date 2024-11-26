package graph

import (
	"testing"
)

func TestProductRecommendation(t *testing.T) {
	// 创建商品图
	pg := NewProductGraph(true, true)

	// 添加商品
	products := []*Product{
		{ID: 1, Labels: []string{"electronics", "phone"}},
		{ID: 2, Labels: []string{"electronics", "laptop"}},
		{ID: 3, Labels: []string{"electronics", "tablet"}},
		{ID: 4, Labels: []string{"furniture", "chair"}},
		{ID: 5, Labels: []string{"furniture", "desk"}},
	}
	for _, product := range products {
		if err := pg.AddProduct(product); err != nil {
			t.Fatalf("Failed to add product: %v", err)
		}
	}

	// 添加相似度
	pg.AddSimilarity(1, 2, 0.8)
	pg.AddSimilarity(1, 3, 0.6)
	pg.AddSimilarity(2, 3, 0.9)
	pg.AddSimilarity(4, 5, 0.7)

	// 测试推荐
	recommendations, err := pg.GetRecommendations(1, 2)
	if err != nil {
		t.Fatalf("Failed to get recommendations: %v", err)
	}

	// 验证推荐结果
	expectedIDs := []int{2, 3}
	for i, rec := range recommendations {
		if rec.ID != expectedIDs[i] {
			t.Errorf("Expected product ID %d, got %d", expectedIDs[i], rec.ID)
		}
	}
}
