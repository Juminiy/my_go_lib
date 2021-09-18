package complicated

import (
	"fmt"
	"github.com/Juminiy/my_go_lib/my_algo/algo_base/data_struct/simple"
	"reflect"
	"strconv"
)

const (
	blankEdge              = 0
	DefaultIsBiDirectional = true
	defaultZero            = 0
	defaultLen             = 1 << 5
	defaultCap             = 1 << maxBin

	nodeIndexNotExist = -1
	nodeNotExist      = 0
	edgeNotExist      = 0
	Epsilon           = "epsilon"
	EdgeEpsilon       = "epsilon"
	EdgeBlank         = "epsilon"
)

// GraphNode equal Value&&i
type GraphNode struct {
	Value interface{}
	i     int
}

func (node *GraphNode) AssignI(_i int) {
	node.i = _i
}
func (node *GraphNode) String() string {
	return "Node[" + strconv.Itoa(node.i) + "]=" + fmt.Sprintf("%v", node.Value)
}

// GraphEdge Value&&i&&j&&k
type GraphEdge struct {
	Value   interface{}
	i, j, k int
}

func (edge *GraphEdge) AssignIJ(_i, _j, _k int) {
	edge.i, edge.j, edge.k = _i, _j, _k
}
func (edge *GraphEdge) String() string {
	return "Edge[" + strconv.Itoa(edge.k) + "](" + strconv.Itoa(edge.i) + " to " + strconv.Itoa(edge.j) + ")=" + fmt.Sprintf("%v", edge.Value)
}
func (graph *AdjGraph) EdgeValueString() string {
	return ""
}

// AdjGraph 如何记录一个点到另一个点的边value edge(i--k-->j)
// Nodes nodeIndex from [0,N-1]
// Edges edgeIndex from [0,M-1]
type AdjGraph struct {
	Nodes    []*GraphNode
	Adjacent map[GraphNode][]*GraphNode
	Edges    []*GraphEdge

	IsUnidirectional   bool
	curNodes, curEdges int // graph.Nodes[nodeIndex].i from 1~N
}

func (graph *AdjGraph) Construct(isUnidirectional bool) {
	graph.Nodes = make([]*GraphNode, 0)
	graph.Edges = make([]*GraphEdge, 0)
	graph.Adjacent = make(map[GraphNode][]*GraphNode, 0)
	graph.IsUnidirectional = isUnidirectional
	graph.curNodes, graph.curEdges = nodeNotExist, edgeNotExist
}

// AddNode i
// nodeIndex = n-1 node.i = n
func (graph *AdjGraph) AddNode(node *GraphNode) {
	if node == nil {
		return
	}
	i := graph.ExistNodeValue(node)
	if i != nodeNotExist {
		return
	}
	graph.curNodes++
	node.i = graph.curNodes
	graph.Nodes = append(graph.Nodes, node)
}

// AddEdge (i--k-->j)
func (graph *AdjGraph) AddEdge(nodeI, nodeJ *GraphNode, edge *GraphEdge) {
	if edge == nil {
		return
	}
	i, j := graph.ExistNodeValue(nodeI), graph.ExistNodeValue(nodeJ)
	_, _, ek := graph.ExistEdge(i, j, edge)
	if i != nodeNotExist && j != nodeNotExist && ek != edgeNotExist {
		graph.Edges[ek-1].Value = edge.Value
		return
	}
	if i == nodeNotExist {
		graph.AddNode(nodeI)
		i = graph.curNodes
	}
	if j == nodeNotExist {
		graph.AddNode(nodeJ)
		j = graph.curNodes
	}
	graph.Adjacent[*nodeI] = append(graph.Adjacent[*nodeI], nodeJ)
	graph.curEdges++
	edge.i, edge.j, edge.k = i, j, graph.curEdges
	graph.Edges = append(graph.Edges, edge)
	if !graph.IsUnidirectional {
		graph.Adjacent[*nodeJ] = append(graph.Adjacent[*nodeJ], nodeI)
		graph.curEdges++
		edge.j, edge.i, edge.k = i, j, graph.curEdges
		graph.Edges = append(graph.Edges, edge)
	}
}

// ExistNodeValue 存在相同值的节点
func (graph *AdjGraph) ExistNodeValue(cNode *GraphNode) int {
	if cNode != nil {
		for _, node := range graph.Nodes {
			// log.Println("node.Value = ",node.Value,"cNode.Value = ",cNode)
			if reflect.DeepEqual(node.Value, cNode.Value) {
				return node.i
			}
		}
	}
	return nodeNotExist
}

// NodeIToNodeIndex nodeIndex=graph.Nodes[nodeIndex].i 求nodeIndex
// 通过node.i 求出 nodeIndex
func (graph *AdjGraph) NodeIToNodeIndex(i int) int {
	for nodeIndex, node := range graph.Nodes {
		if node.i == i {
			return nodeIndex
		}
	}
	return nodeIndexNotExist
}

func (graph *AdjGraph) TestINodeIndex() {
	for nodeIndex, node := range graph.Nodes {
		fmt.Printf("graph.Nodes[%d] nodeIndex is %d,node.i is %d\n", nodeIndex, graph.NodeIToNodeIndex(node.i), node.i)
	}
}
func (graph *AdjGraph) ExistEdgeValue(cEdge *GraphEdge) bool {
	if cEdge != nil {
		for _, edge := range graph.Edges {
			if reflect.DeepEqual(edge.Value, cEdge.Value) {
				return true
			}
		}
	}
	return false
}

// ExistEdge 存在相同的边
// 边的i,j传不进来 导致无法调用ei,ej,ek
func (graph *AdjGraph) ExistEdge(i, j int, cEdge *GraphEdge) (int, int, int) {
	if cEdge != nil {
		for _, edge := range graph.Edges {
			if reflect.DeepEqual(edge.Value, cEdge.Value) &&
				edge.i == i &&
				edge.j == j {
				return edge.i, edge.j, edge.k
			}
		}
	}
	return nodeNotExist, nodeNotExist, edgeNotExist
}
func (graph *AdjGraph) Amount() (int, int) {
	return graph.curNodes + 1, graph.curEdges + 1
}
func (graph *AdjGraph) PrintNodes(tAddress []interface{}) {
	// tAddress = tAddress.(*GraphNode)
	for i, value := range tAddress {
		fmt.Printf("Location[%d]'s address is[%p],value is[%v]\n", i, &value, value)
	}
}

func (graph *AdjGraph) WalkFromNodeI(i int, edgeValue interface{}) *MySet {
	if i == nodeNotExist || edgeValue == nil {
		return nil
	}
	resultSet := &MySet{}
	resultSet.Construct()
	for _, edge := range graph.Edges {
		if edge.i == i && reflect.DeepEqual(edgeValue, edge.Value) {
			resultSet.Insert(graph.Nodes[graph.NodeIToNodeIndex(edge.j)])
		}
	}
	return resultSet
}

func (graph *AdjGraph) WalkFromNodeIOnlyEpsilon(i int) *MySet {
	if i == nodeNotExist {
		return nil
	}
	resultSet, visNodes := &MySet{}, &MySet{}
	resultSet.Construct()
	visNodes.Construct()
	walkQueue := &simple.MyQueue{}
	walkQueue.Push(i)
	resultSet.Insert(graph.Nodes[graph.NodeIToNodeIndex(i)])
	for !walkQueue.IsEmpty() {
		inode, _ := walkQueue.Front()
		visNodes.Insert(inode)
		walkQueue.Pop()
		for _, edge := range graph.Edges {
			if edge.i == inode && reflect.DeepEqual(edge.Value, EdgeEpsilon) && !visNodes.Exist(edge.j) {
				resultSet.Insert(graph.Nodes[graph.NodeIToNodeIndex(edge.j)])
				walkQueue.Push(edge.j)
			}
		}
	}
	return resultSet
}

func (graph *AdjGraph) StartWithIndexEdge(nodeIndex int, a interface{}) *MySet {
	mySet := &MySet{}
	mySet.Construct()
	for _, edge := range graph.Edges {
		if edge.i == nodeIndex && edge.Value == a {
			mySet.Insert(edge)
		}
	}
	return mySet
}

// CalculateDiffValueEdge 不带空边
func (graph *AdjGraph) CalculateDiffValueEdge() *MySet {
	if graph == nil {
		return nil
	}
	resultSet := &MySet{}
	resultSet.Construct()
	for _, edge := range graph.Edges {
		if !reflect.DeepEqual(edge.Value, EdgeEpsilon) {
			resultSet.Insert(edge.Value)
		}
	}
	return resultSet
}
func (graph *AdjGraph) CalculateDiffValueNode() *MySet {
	if graph == nil {
		return nil
	}
	resultSet := &MySet{}
	resultSet.Construct()
	for _, node := range graph.Nodes {
		resultSet.Insert(node.Value)
	}
	return resultSet
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

// ConnectedComponents calculate graph Connectivity
// if 1 is connected else not
func (graph *AdjGraph) ConnectedComponents() int {
	return 0
}

func (graph *AdjGraph) FlushSelfToClearMemory() {

}

// IsIsolatedNode ensure graph isolatedNode
// the graph that one node owns only one different value else meaningless
func (graph *AdjGraph) IsIsolatedNode(nodeValue interface{}) {

}
