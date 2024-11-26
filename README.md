# Graph Algorithms in Go

本项目旨在提供一个用 Go 语言实现的图数据结构及其相关算法的教学工具。通过模块化设计，涵盖了基础图操作、遍历算法、路径查找等多种功能。

## 功能列表
- 支持有向图和无向图
- 支持带权图和无权图
- 图的遍历算法：
  - 广度优先遍历 (BFS)
  - 深度优先遍历 (DFS)
  - 随机游走 (Random Walk)
- 路径查找：
  - Dijkstra 最短路径
  - Floyd-Warshall 多源最短路径
- 图特性检测：
  - 是否为无环图 (DAG)
- 可视化支持：
  - PlantUML 格式导出



## 快速开始
```go
package main

import (
    "fmt"
    "pkg/graph"
)

func main() {
    g := graph.NewAdjacencyList(true, true) // 有向加权图
    g.AddNode(1)
    g.AddNode(2)
    g.AddEdge(1, 2, 3.5)

    distances, _, _ := graph.Dijkstra(g, 1)
    fmt.Println("最短路径距离:", distances)
}
```
