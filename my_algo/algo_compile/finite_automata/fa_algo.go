package finite_automata

import (
	ds "github.com/Juminiy/my_go_lib/my_algo/algo_base/data_struct/complicated"
	"github.com/Juminiy/my_go_lib/my_algo/algo_base/data_struct/simple"
	"github.com/Juminiy/my_go_lib/my_algo/algo_compile/input_struct"
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

func ConstructGraph(inputArr []input_struct.EdgeInput) *ds.AdjGraph {
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

// NFAGenerateSubsetsAndConstructDFA 为了消除 epsilonEdge
func NFAGenerateSubsetsAndConstructDFA(nfaGraph *ds.AdjGraph, startNodes []interface{}) (*ds.MySet, *ds.AdjGraph) {
	C := &ds.MySet{}
	C.Construct() // 幂集合，集合存元素的是单个不重复子集
	startSet := &ISet{}
	startSet.Construct()
	startSet.CharSet.SliceBatchInsert(startNodes)
	edgeValueSet := nfaGraph.CalculateDiffValueEdge()
	T0 := EpsilonClosure(nfaGraph, startSet)
	setQueue := &simple.MyQueue{}
	setQueue.Push(T0) // 因为index会change，所以用队列，队列元素为单个不重复子集
	C.Insert(T0)      // dfa中的每个节点和新的nfa节点完全不同
	dfaGraph := &ds.AdjGraph{}
	dfaGraph.Construct(true) // 幂集中的每个子点集作为新的dfa的一个新节点，将每个节点（子集）rename 0|1|2|...|n
	dfaNodeValue := nodeZeroInt
	dfaNodeQueue := &simple.MyQueue{}
	dfaNodeQueue.Push(dfaNodeValue)
	ExistNodesMap := make(map[*ISet]int, 0)
	dfaNodeValue++
	for !setQueue.IsEmpty() && !dfaNodeQueue.IsEmpty() {
		Tx, _ := setQueue.Front()
		geneNodeValue, _ := dfaNodeQueue.Front()
		setQueue.Pop()
		dfaNodeQueue.Pop()
		for edgeValue, _ := range edgeValueSet.ImmutableMap {
			tMoveTx := Move(nfaGraph, Tx.(*ISet), edgeValue)
			tMoveTx = SeparateOnlySetValue(tMoveTx)
			TxM := EpsilonClosure(nfaGraph, tMoveTx)
			rateNodeValue := nonNode
			if !C.DeepExist(TxM) {
				C.Insert(TxM)
				ExistNodesMap[TxM] = dfaNodeValue
				setQueue.Push(TxM)
				dfaNodeQueue.Push(dfaNodeValue)
				rateNodeValue = dfaNodeValue
				dfaNodeValue++
			} else {
				rateNodeValue = ExistNodesMap[TxM]
			}
			dfaGraph.AddEdge(&ds.GraphNode{Value: geneNodeValue}, &ds.GraphNode{Value: rateNodeValue}, &ds.GraphEdge{Value: edgeValue}) // 无论是否出现过，都必须加入到边集中去，并rename
		}
	}
	return C, dfaGraph
}

// GenerateSubSets C is union of all subsets
// 求子集依赖于底层的GraphAPI提供支持
func GenerateSubSets(faGraph *ds.AdjGraph, nodes []interface{}) *ds.MySet {
	C, _ := NFAGenerateSubsetsAndConstructDFA(faGraph, nodes)
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

func RegexToNFA(regex string) *ds.AdjGraph {
	return nil
}

func NFAToRegex(faGraph *ds.AdjGraph) string {
	return ""
}

func NFAToDFA(nfa *ds.AdjGraph, startNodes []interface{}) *ds.AdjGraph {
	_, dfa := NFAGenerateSubsetsAndConstructDFA(nfa, startNodes)
	return dfa
}

func MinDFA() *ds.AdjGraph {
	return nil
}
