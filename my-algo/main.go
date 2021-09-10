package main

import (
	"fmt"
	"my_algo/algo_base/algo_basic"
)

func main() {
	sort := &algo_basic.MySort{true}
	n, arr := 0, make([]int, 1e5+1, 1e5+1)
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	sort.MyMergeSort(arr, 0, n-1)
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", arr[i])
	}
}
