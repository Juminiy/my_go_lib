package data_struct

import (
	"errors"
	"sort"
)

type MyHeap struct {
	sort.IntSlice
	isMax bool
}

func (heap *MyHeap) Less(i, j int) bool {
	if heap.isMax {
		return heap.IntSlice[i] > heap.IntSlice[j]
	} else {
		return heap.IntSlice[i] < heap.IntSlice[j]
	}
}
func (heap *MyHeap) IsEmpty() bool {
	return heap.IntSlice.Len() == 0
}
func (heap *MyHeap) Push(value interface{}) {
	heap.IntSlice = append(heap.IntSlice, value.(int))
}
func (heap *MyHeap) Pop() error {
	if heap.IsEmpty() {
		return errors.New("Index out of bounds,heap len = 0! ")
	}
	heap.IntSlice = heap.IntSlice[:len(heap.IntSlice)-1]
	return nil
}
func (heap *MyHeap) Top() (interface{}, error) {
	if heap.IsEmpty() {
		return nil, errors.New("Index out of bounds,heap len = 0! ")
	}
	return heap.IntSlice[len(heap.IntSlice)-1], nil
}
