package algo_basic

import (
	"fmt"
	"testing"
)

func TestMySort_MyQuickSort(t *testing.T) {
	sort := &MySort{true}
	fmt.Println(sort.MyMergeSort([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, 0, 9))
}

func TestMyAbs(t *testing.T) {
	fmt.Println(IntToChar(9))
	fmt.Println(CharToInt('0'))
}
