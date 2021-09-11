package complicated

import (
	"fmt"
	"github.com/Juminiy/my_go_lib/my-algo/algo_base/data_struct/simple"
)

const (
	blankEdge        = 0
	DefaultDirection = true
	defaultZero      = 0
	defaultLen       = 1 << 5
	defaultCap       = 1 << maxBin
)

type GraphNode struct {
	Value interface{}
}
type GraphEdge struct {
	Value interface{}
}
type AdjGraph struct {
	Nodes            []*GraphNode
	Adjacent         map[GraphNode][]*GraphNode
	Edges            []*GraphEdge
	IsUnidirectional bool
}

var seqGraph []interface{}

func (graph *AdjGraph) Construct(isUnidirectional bool) {
	graph.Nodes = make([]*GraphNode, 0)
	graph.Edges = make([]*GraphEdge, 0)
	graph.IsUnidirectional = isUnidirectional
	graph.Adjacent = make(map[GraphNode][]*GraphNode, 0)
}
func (graph *AdjGraph) AddNode(node *GraphNode) {
	graph.Nodes = append(graph.Nodes, node)
}

func (graph *AdjGraph) AddEdge(nodeI, nodeJ *GraphNode, edge *GraphEdge) {

	graph.Adjacent[*nodeI] = append(graph.Adjacent[*nodeI], nodeJ)
	if !graph.IsUnidirectional {
		graph.Adjacent[*nodeJ] = append(graph.Adjacent[*nodeJ], nodeI)
	}
	graph.Edges = append(graph.Edges, edge)
}
func (graph *AdjGraph) PrintNodes(tAddress []interface{}) {
	// tAddress = tAddress.(*GraphNode)
	for i, value := range tAddress {
		fmt.Printf("Location[%d]'s address is[%p],value is[%v]\n", i, &value, value)
	}
}
func (graph *AdjGraph) BfsGraph() []interface{} {
	// fmt.Println(&graph);fmt.Println("---------------------")
	if graph == nil || graph.Nodes == nil || len(graph.Nodes) == 0 {
		return nil
	}
	q, seqNodes, visNode := &simple.MyQueue{}, make([]interface{}, 0), &MySet{}
	visNode.Construct()
	q.Push(graph.Nodes[0])
	for !q.IsEmpty() {
		var curNode *GraphNode
		if tNode, err := q.Front(); err == nil {
			curNode = tNode.(*GraphNode)
			seqNodes = append(seqNodes, curNode)
			visNode.Insert(curNode)
		}
		q.Pop()
		// fmt.Println(tNode)
		// adjNode := graph.Adjacent[*tNode]
		// fmt.Println(adjNode)
		for _, adjNode := range graph.Adjacent[*curNode] {
			if !visNode.Exist(adjNode) {
				q.Push(adjNode)
			}
		}
	}
	return seqNodes
}

// DfsGraph NonRecursion dfs
func (graph *AdjGraph) DfsGraph() []interface{} {
	if graph == nil || graph.Nodes == nil || len(graph.Nodes) == 0 {
		return nil
	}
	s, seqNodes, visNodes := &simple.MyStack{}, make([]interface{}, 0), &MySet{}
	visNodes.Construct()
	curNode := graph.Nodes[0]
	for !s.IsEmpty() || curNode != nil {
		for curNode != nil {
			s.Push(curNode)
			seqNodes = append(seqNodes, curNode)
			visNodes.Insert(curNode)
			if graph.Adjacent[*curNode] != nil {
				curNode = graph.Adjacent[*curNode][0]
				if visNodes.Exist(curNode) {
					curNode = nil
				}
			} else {
				curNode = nil
			}
		}
		if !s.IsEmpty() {
			if tNode, err := s.Top(); err == nil {
				curNode = tNode.(*GraphNode)
			}
			for _, tNode := range graph.Adjacent[*curNode] {
				if !visNodes.Exist(tNode) {
					s.Push(tNode)
					curNode = tNode
					break
				}
			}
			if err := s.Pop(); err == nil {
				if visNodes.Exist(curNode) {
					curNode = nil
				}
			}
		}
	}
	return seqNodes
}
