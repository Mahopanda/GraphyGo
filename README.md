# GraphyGo：學習 Golang 與圖算法的教學工具

該專案用 Go 語言 開發的學習型專案，專注於實現圖（Graph）相關的基礎操作與算法，讓使用者透過實作和實驗來更深入理解圖的結構與應用。

## 適合對象

- 想學 Go 的基礎語法，但又想實作專案的人。
- 想了解圖的資料結構和應用（例如推薦系統、路徑規劃等）。
- 想建立一個基礎框架來開發更複雜的圖算法。

可以從 examples/ 內的範例開始，每個範例都是針對特定問題或演算法。

## 功能列表

- 基本圖操作：
  - 支援 有向圖 和 無向圖
  - 支援 加權圖 和 無權圖
  - 圖結構的可視化輸出（支援 PlantUML）
- 遍歷方法：
  - 廣度優先搜尋 (BFS)
  - 深度優先搜尋 (DFS)
  - 隨機遊走 (Random Walk)
- 路徑查找：
  - Dijkstra 最短路徑算法
  - A 啟發式搜索\*
- 其他進階功能：
  - 拓撲排序：解決任務依賴問題（如課程安排）
  - DAG 檢測：判斷是否為無環圖
  - 團（Clique）查找：探索高連接子圖
  - 相似性推薦：基於圖的商品推薦系統

## 開發目標

- 學習 Golang 的基本語法和特性。
- 了解和練習圖的基本概念與演算法。
- 透過實作，加深對圖在實際應用中的理解（例如：路徑規劃、推薦系統）。
- 這套工具目標是教學用途，設計上會盡量簡單明瞭，方便快速上手。

## 如何開始使用

### 安裝方式

1. 下載專案：

   ```bash
   git clone https://github.com/Mahopanda/GraphyGo.git
   ```

2. 安裝必要的套件： 進入專案目錄後執行：

   ```bash
   go mod tidy
   ```

3. 開始開發或學習： 可以直接修改 examples/ 資料夾內的範例程式，也可以寫自己的小測試程式。

### 範例教學

以下範例展示如何創建一個圖、添加節點與邊，並進行基本的遍歷和演算法操作。

#### 範例一：基本圖操作

```go
package main

import (
"fmt"
"github.com/Mahopanda/GraphyGo/pkg/graph"
)

func main() {
// 創建一個無向無權圖
g := graph.NewAdjacencyList(false, false)

    // 添加節點
    nodes := []int{1, 2, 3, 4}
    for _, node := range nodes {
    	g.AddNode(node)
    }

    // 添加邊
    g.AddEdge(1, 2, 0)
    g.AddEdge(2, 3, 0)
    g.AddEdge(3, 4, 0)
    g.AddEdge(4, 1, 0)

    // 輸出節點與邊的資訊
    fmt.Println("節點數量:", g.NodeCount())
    fmt.Println("邊的數量:", g.EdgeCount())
    fmt.Println("邊 1->2 是否存在:", g.HasEdge(1, 2))

    // 圖的可視化輸出
    uml, _ := graph.ToPlantUML(g)
    fmt.Println("\nPlantUML 輸出:")
    fmt.Println(uml)

}
```

#### 範例二：最短路徑計算（Dijkstra）

```go
package main

import (
"fmt"
"github.com/Mahopanda/GraphyGo/pkg/graph"
)

func main() {
// 創建一個有向加權圖
g := graph.NewAdjacencyList(true, true)

    // 添加節點與邊
    g.AddNode(0)
    g.AddNode(1)
    g.AddNode(2)
    g.AddEdge(0, 1, 1)
    g.AddEdge(1, 2, 2)

    // 使用 Dijkstra 算法計算最短路徑
    distances, predecessors, _ := graph.Dijkstra(g, 0)

    // 輸出結果
    fmt.Println("最短路徑距離:", distances)
    fmt.Println("前驅節點:", predecessors)

}
```

### 專案結構

核心功能模組：

- adjacency_list.go：提供圖的基本操作（新增節點、添加邊、獲取鄰居等）。
- traversal.go：實現 BFS、DFS 與隨機遊走。
- shortest_path.go：實現 Dijkstra。
- iterator.go：圖的迭代器（例如：拓撲排序）。
- product_graph.go：實現商品圖與推薦功能。
- plantuml.go：圖的可視化輸出。

測試與範例：

- examples/：範例程式，展示如何使用 GraphyGo 的功能。
- 單元測試涵蓋部份功能。
