package service

import (
	"github.com/Juminiy/my_go_lib/my-algo/algo_base/algo_basic"
	"strconv"
)

const (
	RecursiveLevel = 31
)

var sort = new(algo_basic.MySort)

func IntMinAlgoService(compValue, compedValue string) int {
	intComp, _ := strconv.Atoi(compValue)
	intComped, _ := strconv.Atoi(compedValue)
	return algo_basic.MinValue(intComp, intComped)
}
func IntQuickSortService(arr []int) []int {
	sort.MyQuickSort(arr, 0, len(arr)-1)
	return arr
}
func IntMergeSortService(arr []int) [][]int {
	sort.IsAsc = true
	return sort.MyMergeSort(arr, 0, len(arr)-1)
}
