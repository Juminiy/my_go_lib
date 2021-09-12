package fa

import (
	"fmt"
	"github.com/Juminiy/my_go_lib/my-algo/algo_base/data_struct/complicated"
	"testing"
)

func TestEpsilonClosure(t *testing.T) {
	adj := &complicated.AdjGraph{}
	adj.Construct(true)
	adj.AddEdge(&complicated.GraphNode{Value: 0}, &complicated.GraphNode{Value: 1}, &complicated.GraphEdge{Value: "epsilon"})
	adj.AddEdge(&complicated.GraphNode{Value: 0}, &complicated.GraphNode{Value: 7}, &complicated.GraphEdge{Value: "epsilon"})
	adj.AddEdge(&complicated.GraphNode{Value: 1}, &complicated.GraphNode{Value: 2}, &complicated.GraphEdge{Value: "epsilon"})
	adj.AddEdge(&complicated.GraphNode{Value: 1}, &complicated.GraphNode{Value: 4}, &complicated.GraphEdge{Value: "epsilon"})
	adj.AddEdge(&complicated.GraphNode{Value: 2}, &complicated.GraphNode{Value: 3}, &complicated.GraphEdge{Value: "a"})
	adj.AddEdge(&complicated.GraphNode{Value: 4}, &complicated.GraphNode{Value: 5}, &complicated.GraphEdge{Value: "b"})
	adj.AddEdge(&complicated.GraphNode{Value: 3}, &complicated.GraphNode{Value: 6}, &complicated.GraphEdge{Value: "epsilon"})
	adj.AddEdge(&complicated.GraphNode{Value: 5}, &complicated.GraphNode{Value: 6}, &complicated.GraphEdge{Value: "epsilon"})
	adj.AddEdge(&complicated.GraphNode{Value: 6}, &complicated.GraphNode{Value: 1}, &complicated.GraphEdge{Value: "epsilon"})
	adj.AddEdge(&complicated.GraphNode{Value: 6}, &complicated.GraphNode{Value: 7}, &complicated.GraphEdge{Value: "epsilon"})
	adj.AddEdge(&complicated.GraphNode{Value: 7}, &complicated.GraphNode{Value: 8}, &complicated.GraphEdge{Value: "a"})
	adj.AddEdge(&complicated.GraphNode{Value: 8}, &complicated.GraphNode{Value: 9}, &complicated.GraphEdge{Value: "b"})
	adj.AddEdge(&complicated.GraphNode{Value: 9}, &complicated.GraphNode{Value: 10}, &complicated.GraphEdge{Value: "b"})
	tSet := &ISet{}
	tSet.Construct()
	tSet.CharSet.Insert(0)
	fmt.Println(Move(adj, tSet, "epsilon"))

}
