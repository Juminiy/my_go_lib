package service

import (
	"github.com/Juminiy/my_go_lib/my-algo/algo_base/algo_basic"
	"github.com/Juminiy/my_go_lib/my-algo/algo_base/data_struct/complicated"
	"github.com/Juminiy/my_go_lib/my-algo/algo_compile/fa"
	"github.com/Juminiy/my_go_lib/my-algo/algo_compile/struc"
	"strconv"
)

const (
	RecursiveLevel = 31
)

var sort = new(algo_basic.MySort)
var zeroSet = &fa.ISet{CharSet: &complicated.MySet{ImmutableMap: map[interface{}]bool{"0": true}}}

func IntMinAlgoService(compValue, compedValue string) int {
	intComp, _ := strconv.Atoi(compValue)
	intComped, _ := strconv.Atoi(compedValue)
	return algo_basic.MinValue(intComp, intComped)
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
func EpsilonClosureService(inputArr []struc.EdgeInput) []interface{} {
	adj := fa.ConstructGraph(inputArr)
	tSet := fa.EpsilonClosure(adj, zeroSet)
	return tSet.CharSet.Slice()
}
