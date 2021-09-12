package fa

import (
	ds "github.com/Juminiy/my_go_lib/my-algo/algo_base/data_struct/complicated"
	"github.com/Juminiy/my_go_lib/my-algo/algo_compile/struc"
)

const (
	epsilon    = "epsilon"
	charQ      = 'q'
	initialNum = '0'

	nonEdge = 0
	nonNode = 0
)

/*
 * function run depends on the graph
 */

func ConstructGraph(inputArr []struc.EdgeInput) *ds.AdjGraph {
	adj := &ds.AdjGraph{}
	adj.Construct(true)
	for _, edge := range inputArr {
		adj.AddEdge(&ds.GraphNode{Value: edge.NodeIValue}, &ds.GraphNode{Value: edge.NodeJValue}, &ds.GraphEdge{Value: edge.EdgeValue})
	}
	return adj
}

// Move 集合I的所有状态经过一次a边到达的状态
func Move(faGraph *ds.AdjGraph, I *ISet, a interface{}) *ISet {
	iSet := &ISet{}
	iSet.Construct()
	if faGraph.ExistEdgeValue(&ds.GraphEdge{Value: a}) {
		for state, _ := range I.CharSet.ImmutableMap { //key value写反了
			nodeIndex := faGraph.ExistNodeValue(&ds.GraphNode{Value: state})
			if a == epsilon {
				iSet.CharSet.Unite(faGraph.WalkFromNodeIndexOnlyEpsilon(nodeIndex))
			} else {
				iSet.CharSet.Unite(faGraph.WalkFromNodeIndex(nodeIndex, a))
			}
		}
	}
	// fmt.Println(iSet)
	return iSet
}

// EpsilonClosure 集合I的所有状态经若干次空边到达的状态 并上 集合I
func EpsilonClosure(faGraph *ds.AdjGraph, I *ISet) *ISet {
	if I == nil || I.CharSet.Len() == 0 {
		return nil
	}
	iSet := &ISet{}
	iSet.Construct()
	for _, state := range I.CharSet.ImmutableMap {
		nodeIndex := faGraph.ExistNodeValue(&ds.GraphNode{Value: state})
		iSet.CharSet.Unite(faGraph.WalkFromNodeIndexOnlyEpsilon(nodeIndex))
	}
	return &ISet{iSet.CharSet.Unite(I.CharSet)}
}

// GenerateSubSets C is union of all subsets
func GenerateSubSets(faGraph *ds.AdjGraph) *ISet {
	C := &ds.MySet{}
	C.Construct()
	return nil
}

func RegexToGraph(regex string) *ds.AdjGraph {
	if len(regex) == 0 {
		return nil
	}
	adj := &ds.AdjGraph{}
	adj.Construct(true)
	startNode, q0Node := &ds.GraphNode{Value: "start"}, &ds.GraphNode{Value: charQ + initialNum}
	adj.AddNode(startNode)
	adj.AddNode(q0Node)
	// subStr := ""
	for _, ch := range regex {
		switch ch {
		case '(':
			{

				break
			}
		case ')':
			{
				break
			}
		case '|':
			{
				break
			}

		default:
			{

			}
		}
	}
	return adj
}
