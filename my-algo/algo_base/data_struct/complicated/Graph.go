package complicated

import (
	"fmt"
	"github.com/Juminiy/my_go_lib/my-algo/algo_base/data_struct/simple"
	"reflect"
	"strconv"
)

const (
	blankEdge        = 0
	DefaultDirection = true
	defaultZero      = 0
	defaultLen       = 1 << 5
	defaultCap       = 1 << maxBin

	nodeNotExist = 0
	edgeNotExist = 0
	EdgeEpsilon  = "epsilon"
)

var curNode, curEdge = nodeNotExist, edgeNotExist

// GraphNode Value&&i
type GraphNode struct {
	Value interface{}
	i     int
}

func (node *GraphNode) AssignI(_i int) {
	node.i = _i
}
func (node *GraphNode) String() string {
	return "Node[" + strconv.Itoa(node.i) + "]=" + node.Value.(string)
}

// GraphEdge Value&&i&&j&&k
type GraphEdge struct {
	Value   interface{}
	i, j, k int
}

func (edge *GraphEdge) AssignIJ(_i, _j int) {
	edge.i, edge.j = _i, _j
}
func (edge *GraphEdge) String() string {
	return "Edge[" + strconv.Itoa(edge.k) + "](" + strconv.Itoa(edge.i) + " to " + strconv.Itoa(edge.j) + ")=" + edge.Value.(string)
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

// AddNode i
func (graph *AdjGraph) AddNode(node *GraphNode) {
	if node == nil {
		return
	}
	i := graph.ExistNodeValue(node)
	if i != nodeNotExist {
		return
	}
	curNode++
	node.i = curNode
	graph.Nodes = append(graph.Nodes, node)
}

// AddEdge (i-k>j)
func (graph *AdjGraph) AddEdge(nodeI, nodeJ *GraphNode, edge *GraphEdge) {
	if edge == nil {
		return
	}
	i, j := graph.ExistNodeValue(nodeI), graph.ExistNodeValue(nodeJ)
	ei, ej, ek := graph.ExistEdge(edge)
	if ek != edgeNotExist {
		return
	}
	if i != nodeNotExist && j != nodeNotExist && ei == i && ej == j {
		graph.Edges[ek].Value = edge.Value
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
	edge.i, edge.j, edge.k = i, j, curEdge
	graph.Edges = append(graph.Edges, edge)
}

// ExistNodeValue 存在相同值的节点
func (graph *AdjGraph) ExistNodeValue(cNode *GraphNode) int {
	if cNode != nil {
		for _, node := range graph.Nodes {
			if reflect.DeepEqual(node.Value, cNode.Value) {
				return node.i
			}
		}
	}
	return nodeNotExist
}

// ExistEdge 存在相同的边
func (graph *AdjGraph) ExistEdge(cEdge *GraphEdge) (int, int, int) {
	if cEdge != nil {
		for _, edge := range graph.Edges {
			if reflect.DeepEqual(edge.Value, cEdge.Value) &&
				edge.i == cEdge.i &&
				edge.j == cEdge.j {
				return edge.i, edge.j, edge.k
			}
		}
	}
	return nodeNotExist, nodeNotExist, edgeNotExist
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
