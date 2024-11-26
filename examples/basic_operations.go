package main

import (
	"fmt"
	"os"

	"github.com/Mahopanda/GraphyGo/pkg/graph"
)

// 基本图操作
func main() {
	// 创建一个无向无权图
	g := graph.NewAdjacencyList(false, false)

	// 添加节点
	nodes := []int{1, 2, 3, 4}
	for _, node := range nodes {
		if err := g.AddNode(node); err != nil {
			fmt.Printf("添加節點失敗: %v\n", err)
		}
	}

	// 添加边
	g.AddEdge(1, 2, 0)
	g.AddEdge(2, 3, 0)
	g.AddEdge(3, 4, 0)
	g.AddEdge(4, 1, 0)

	// 打印节点和边的数量
	fmt.Println("节点数量:", g.NodeCount()) // 输出: 4
	fmt.Println("边数量:", g.EdgeCount())  // 输出: 4

	// 检查边是否存在
	fmt.Println("边 1->2 是否存在:", g.HasEdge(1, 2)) // 输出: true
	fmt.Println("边 3->1 是否存在:", g.HasEdge(3, 1)) // 输出: false

	uml, err := graph.ToPlantUML(g)
	if err != nil {
		fmt.Printf("生成 PlantUML 失败: %v\n", err)
		return
	}
	// 将 PlantUML 输出保存到文件
	err = os.WriteFile("graph.puml", []byte(uml), 0644)
	if err != nil {
		fmt.Printf("保存 PlantUML 文件失败: %v\n", err)
		return
	}
	fmt.Println("\n已生成 PlantUML 文件 'graph.puml'，可使用在线工具查看图形。")

}
