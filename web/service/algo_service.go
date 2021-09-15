package service

import (
	"github.com/Juminiy/my_go_lib/my_algo/algo_base/algo_basic"
	"github.com/Juminiy/my_go_lib/my_algo/algo_base/data_struct/complicated"
	"github.com/Juminiy/my_go_lib/my_algo/algo_compile/finite_automata"
	"github.com/Juminiy/my_go_lib/my_algo/algo_compile/input_struct"
	"log"
	"strconv"
)

const (
	RecursiveLevel = 31
)

var sort = new(algo_basic.MySort)
var zeroSet = &finite_automata.ISet{CharSet: &complicated.MySet{ImmutableMap: map[interface{}]bool{0: true}}}

func IntMinAlgoService(compValue, compedValue string) int {
	intComp, _ := strconv.Atoi(compValue)
	intComped, _ := strconv.Atoi(compedValue)
	return algo_basic.MinValue(intComp, intComped)
}
func IntMaxAlgoService(compValue, compedValue string) int {
	intComp, _ := strconv.Atoi(compValue)
	intComped, _ := strconv.Atoi(compedValue)
	return algo_basic.MaxValue(intComp, intComped)
}
func IntQuickSortService(arr []int) []int {
	sort.IsAsc = true
	sort.MyQuickSort(arr, 0, len(arr)-1)
	return arr
}
func IntMergeSortService(arr []int) [][]int {
	sort.IsAsc = true
	return sort.MyMergeSort(arr, 0, len(arr)-1)
}
func EpsilonClosureService(edges []input_struct.EdgeInput, nodes []interface{}) []interface{} {
	adj := finite_automata.ConstructGraph(edges)
	//log.Println(adj.Nodes)
	if nodes == nil {
		nodes = zeroSet.CharSet.Slice()
	}
	tSet := finite_automata.EpsilonClosure(adj, &finite_automata.ISet{CharSet: complicated.SliceToSet(nodes)})
	return tSet.CharSet.Slice()
}

func ConstructSubSetService(edges []input_struct.EdgeInput, nodes []interface{}) string {
	adj := finite_automata.ConstructGraph(edges)
	if nodes == nil {
		nodes = zeroSet.CharSet.Slice()
	}
	// log.Println(adj.Nodes,adj.Edges)
	tSet := finite_automata.GenerateSubSets(adj, nodes)
	log.Println(tSet)
	return tSet.String()
}
