package algo_leetcode

import (
	"fmt"
	"testing"
)

func TestMin(t *testing.T) {
	//fmt.Println(coinChange([]int{2},3))
	x := singleNumber([]int{1, 2, 1, 3, 2, 5})
	fmt.Println(x, x&-x)
}
