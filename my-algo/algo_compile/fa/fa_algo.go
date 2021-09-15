package fa

import (
	ds "github.com/Juminiy/my_go_lib/my-algo/algo_base/data_struct/complicated"
	"github.com/Juminiy/my_go_lib/my-algo/algo_base/data_struct/simple"
	"github.com/Juminiy/my_go_lib/my-algo/algo_compile/struc"
)

// 多线程环境线程常量不可用 
const (
	epsilon      = "epsilon"
	charQ        = 'q'
	initialNum   = '0'
	nodeZeroInt  = 0
	nodeZeroChar = '0'
	nonEdge      = 0
	nonNode      = 0
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
	if I == nil || I.CharSet.Len() == 0 {
		return nil
	}
	iSet := &ISet{}
	iSet.Construct()
	if faGraph.ExistEdgeValue(&ds.GraphEdge{Value: a}) {
		for state, _ := range I.CharSet.ImmutableMap { //key value写反了
			nodeIndex := faGraph.ExistNodeValue(&ds.GraphNode{Value: state})
			if a == epsilon {
				iSet.CharSet = iSet.CharSet.Unite(faGraph.WalkFromNodeIOnlyEpsilon(nodeIndex))
			} else {
				iSet.CharSet = iSet.CharSet.Unite(faGraph.WalkFromNodeI(nodeIndex, a))
			}
		}
	}
	return iSet
}

// EpsilonClosure 集合I的所有状态经若干次空边到达的状态 并上 集合I
func EpsilonClosure(faGraph *ds.AdjGraph, I *ISet) *ISet {
	if I == nil || I.CharSet.Len() == 0 {
		return nil
	}
	iSet := &ISet{CharSet: I.CharSet}
	// log.Println(faGraph.Nodes)
	for state, _ := range I.CharSet.ImmutableMap {
		nodeI := faGraph.ExistNodeValue(&ds.GraphNode{Value: state})
		// log.Println("nodeI = ",nodeI)
		edgeEpsilonNodes := faGraph.WalkFromNodeIOnlyEpsilon(nodeI)
		iSet.CharSet = iSet.CharSet.Unite(NodeSetToIntValueSet(&ISet{CharSet: edgeEpsilonNodes}))
	}
	return iSet
}

func NodeSetToIntValueSet(I *ISet) *ds.MySet {
	if I == nil || I.CharSet.Len() <= 0 {
		return nil
	}
	valueSet := &ds.MySet{}
	valueSet.Construct()
	for value, _ := range I.CharSet.ImmutableMap {
		valueSet.Insert(value.(*ds.GraphNode).Value)
	}
	return valueSet
}

// GenerateSubSets C is union of all subsets
// 求子集依赖于底层的GraphAPI提供支持
func GenerateSubSets(faGraph *ds.AdjGraph, nodes []interface{}) *ds.MySet {
	C := &ds.MySet{} // 幂集合，集合存元素的是单个不重复子集
	C.Construct()    // 因为序index会change，所以用队列，队列元素为单个不重复子集
	startSet := &ISet{}
	startSet.Construct()
	startSet.CharSet.SliceBatchInsert(nodes)
	edgeValueSet := faGraph.CalculateDiffValueEdge()
	T0 := EpsilonClosure(faGraph, startSet)
	setQueue := &simple.MyQueue{}
	setQueue.Push(T0)
	C.Insert(T0)
	for !setQueue.IsEmpty() {
		Tx, _ := setQueue.Front()
		setQueue.Pop()
		for edgeValue, _ := range edgeValueSet.ImmutableMap {
			tMoveTx := Move(faGraph, Tx.(*ISet), edgeValue)
			tMoveTx = SeparateOnlySetValue(tMoveTx)
			TxM := EpsilonClosure(faGraph, tMoveTx)
			if !C.DeepExist(TxM) {
				C.Insert(TxM)
				setQueue.Push(TxM)
			}
		}
	}
	return C
}

//
//func TwoDimSetToTwoDimSlice(C *ds.MySet) [][]interface{} {
//	if C == nil || C.Len() < 0 {
//		return nil
//	}
//	Dim2Slice := make([][]interface{},0,0)
//	for eleArr,_ := range C.ImmutableMap {
//
//		Dim2Slice = append(Dim2Slice, eleArr)
//	}
//	return Dim2Slice
//}

func SubSetByOrder(C *ds.MySet) *ds.MySet {
	if C == nil || C.Len() == 0 {
		return nil
	}
	for subSet, _ := range C.ImmutableMap {
		C.Insert(subSet.(*ISet).CharSet.SortSetToSlice())
	}
	return C
}
func SeparateOnlySetValue(I *ISet) *ISet {
	if I == nil || I.CharSet.Len() == 0 {
		return nil
	}
	iSet := &ISet{}
	iSet.Construct()
	for ele, _ := range I.CharSet.ImmutableMap {
		iSet.CharSet.Insert(ele.(*ds.GraphNode).Value)
	}
	return iSet
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
