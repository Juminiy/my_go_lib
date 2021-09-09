package st

import (
	"fmt"
	"testing"
)

func TestST_QueryRange(t *testing.T) {
	arrLen, queryTimes, l, r := 0, 0, 0, 0
	fmt.Scan(&arrLen)
	arr := make([]int, arrLen, arrLen)
	for i := 0; i < arrLen; i++ {
		fmt.Scan(&arr[i])
	}
	stList := &ST{}
	stList.ConstructST(arr)
	for i := 0; i < queryTimes; i++ {
		fmt.Scan(&l, &r)
		fmt.Println(stList.QueryRange(l, r))
	}
}

func TestName(t *testing.T) {

}
