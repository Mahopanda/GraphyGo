package main

import (
	"fmt"
	"os"

	"github.com/Mahopanda/GraphyGo/pkg/graph"
)

// 1
// 0 --------→ 1
//  \          ↓
//   \    4    2
//    ↘        ↓
// 	 2 ----→ 3
// 		 1

// 最短路径算法
func main() {
	// 创建一个带权图
	g := graph.NewAdjacencyList(true, true)

	// 添加节点
	g.AddNode(0)
	g.AddNode(1)
	g.AddNode(2)
	g.AddNode(3)

	// 添加双向边，使图更有趣
	g.AddEdge(0, 1, 1)
	g.AddEdge(1, 0, 1) // 添加反向边
	g.AddEdge(1, 2, 2)
	g.AddEdge(2, 1, 2) // 添加反向边
	g.AddEdge(0, 2, 4)
	g.AddEdge(2, 0, 4) // 添加反向边
	g.AddEdge(2, 3, 1)
	g.AddEdge(3, 2, 1) // 添加反向边

	// 使用 Dijkstra 算法查找从节点 0 出发的最短路径
	distances, predecessors, _ := graph.Dijkstra(g, 0)

	// 输出距离
	fmt.Println("最短路径距离:", distances) // 输出: map[0:0 1:1 2:3 3:4]

	// 输出前驱节点
	fmt.Println("前驱节点:", predecessors) // 输出: map[1:0 2:1 3:2]

	// 输出图的可视化
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
