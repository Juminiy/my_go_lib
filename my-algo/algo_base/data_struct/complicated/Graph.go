package complicated

import "github.com/Juminiy/my_go_lib/my-algo/algo_base/data_struct/simple"

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

func (graph *AdjGraph) BfsGraph() []interface{} {
	if graph == nil {
		return nil
	}
	q, seqNode, visNode := &simple.MyQueue{}, make([]interface{}, 0), &MySet{}
	visNode.Construct()
	q.Push(graph.Nodes[0])
	for !q.IsEmpty() {
		var tNode *GraphNode
		if tNode, err := q.Front(); err == nil {
			tNode = tNode.(*GraphNode)
			seqNode = append(seqNode, tNode)
			visNode.Insert(tNode)
		}
		q.Pop()
		for adjNode := range graph.Adjacent[*tNode] {
			if !visNode.Exist(adjNode) {
				q.Push(adjNode)
			}
		}
	}
	return seqNode
}

func DfsGraph(graph *AdjGraph) []interface{} {
	return nil
}
