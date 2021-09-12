package complicated

import (
	"fmt"
	"github.com/Juminiy/my_go_lib/my-algo/algo_base/data_struct/simple"
	"reflect"
)

const (
	blankEdge        = 0
	DefaultDirection = true
	defaultZero      = 0
	defaultLen       = 1 << 5
	defaultCap       = 1 << maxBin

	nodeNotExist = 0
	edgeNotExist = 0
)

var curNode, curEdge = nodeNotExist, edgeNotExist

type GraphNode struct {
	Value interface{}
	I     int
}
type GraphEdge struct {
	Value   interface{}
	I, J, K int
}

// AdjGraph 如何记录一个点到另一个点的边value
type AdjGraph struct {
	Nodes            []*GraphNode
	Adjacent         map[GraphNode][]*GraphNode
	Edges            []*GraphEdge
	IsUnidirectional bool
}

func (graph *AdjGraph) Construct(isUnidirectional bool) {
	graph.Nodes = make([]*GraphNode, 0)
	graph.Edges = make([]*GraphEdge, 0)
	graph.IsUnidirectional = isUnidirectional
	graph.Adjacent = make(map[GraphNode][]*GraphNode, 0)
}
func (graph *AdjGraph) AddNode(node *GraphNode) {
	curNode, node.I = curNode+1, curNode
	graph.Nodes = append(graph.Nodes, node)
}

func (graph *AdjGraph) AddEdge(nodeI, nodeJ *GraphNode, edge *GraphEdge) {
	i, j, k := graph.ExistNodeValue(nodeI), graph.ExistNodeValue(nodeJ), graph.DeepExistEdge(nodeI, nodeJ, edge)
	if k != edgeNotExist {
		return
	}
	if i == nodeNotExist {
		graph.AddNode(nodeI)
		i = curNode
	}
	if j == nodeNotExist {
		graph.AddNode(nodeJ)
		j = curNode
	}
	graph.Adjacent[*nodeI] = append(graph.Adjacent[*nodeI], nodeJ)
	curEdge++
	if !graph.IsUnidirectional {
		graph.Adjacent[*nodeJ] = append(graph.Adjacent[*nodeJ], nodeI)
		curEdge++
	}
	edge.I, edge.J, edge.K = i, j, k
	graph.Edges = append(graph.Edges, edge)
}

func (graph *AdjGraph) ExistNodeValue(cNode *GraphNode) int {
	for _, node := range graph.Nodes {
		if reflect.DeepEqual(node.Value, cNode.Value) {
			return node.I
		}
	}
	return nodeNotExist
}
func (graph *AdjGraph) ExistEdgeValue(cEdge *GraphEdge) (int, int) {
	for _, edge := range graph.Edges {
		if reflect.DeepEqual(edge.Value, cEdge.Value) {
			return edge.I, edge.J
		}
	}
	return edgeNotExist, edgeNotExist
}
func (graph *AdjGraph) DeepExistEdge(nodeI, nodeJ *GraphNode, cEdge *GraphEdge) int {
	i, j := graph.ExistEdgeValue(cEdge)
	if graph.ExistNodeValue(nodeI) != nodeNotExist &&
		graph.ExistNodeValue(nodeJ) != nodeNotExist &&
		i != edgeNotExist && j != edgeNotExist {
		for _, node := range graph.Adjacent[*nodeI] {
			if reflect.DeepEqual(node.Value, nodeJ.Value) {
				return edgeNotExist
			}
		}
		return edgeNotExist
	}
	return edgeNotExist
}
func (graph *AdjGraph) Amount() (int, int) {
	return curNode, curEdge
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
	for _, curNode := range graph.Nodes {
		if !visNode.Exist(curNode) {
			q.Push(curNode)
		}
		for !q.IsEmpty() {
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
	for _, curNode := range graph.Nodes {
		for !s.IsEmpty() || (curNode != nil && !visNodes.Exist(curNode)) {
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
	}

	return seqNodes
}
