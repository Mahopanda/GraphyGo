package main

import (
	"fmt"

	"github.com/Mahopanda/GraphyGo/pkg/graph"
)

// Clique（团）是图中的一个完全子图，其中每两个节点都直接相连。通过寻找团，可以识别网络中高度连接的子群体（如社交网络中的紧密社交圈）。

// CliqueIterator 查找所有大小為指定值的團
func CliqueIterator(g *graph.AdjacencyList, size int) [][]int {
	cliques := [][]int{}
	nodes := []int{}

	// 將所有節點存入列表
	for node := range g.GetNodes() {
		nodes = append(nodes, node)
	}

	// 遞歸搜索團
	var search func(path []int, candidates []int)
	search = func(path []int, candidates []int) {
		if len(path) == size {
			cliques = append(cliques, append([]int(nil), path...))
			return
		}

		for i, node := range candidates {
			// 檢查是否與目前路徑中的所有節點相連
			valid := true
			for _, p := range path {
				if !g.HasEdge(p, node) {
					valid = false
					break
				}
			}
			if valid {
				// 繼續搜索
				search(append(path, node), candidates[i+1:])
			}
		}
	}

	search([]int{}, nodes)
	return cliques
}

func main() {
	// 創建一個無向圖
	g := graph.NewAdjacencyList(false, false)

	// 添加節點（用戶編號）
	users := []int{1, 2, 3, 4, 5}
	for _, user := range users {
		g.AddNode(user)
	}

	// 添加邊（用戶之間的社交關係）
	connections := []struct {
		from, to int
	}{
		{1, 2}, {1, 3}, {2, 3}, // 團: 1-2-3
		{3, 4}, {4, 5}, // 團: 3-4-5
	}
	for _, edge := range connections {
		g.AddEdge(edge.from, edge.to, 0)
		g.AddEdge(edge.to, edge.from, 0) // 無向圖需要雙向添加
	}

	// 查找大小為 3 的團
	cliques := CliqueIterator(g, 3)

	// 輸出結果
	fmt.Println("大小為 3 的團:")
	for _, clique := range cliques {
		fmt.Println(clique)
	}
}
