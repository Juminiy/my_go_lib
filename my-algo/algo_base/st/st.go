package st

import "github.com/Juminiy/my_go_lib/my-algo/algo_base/algo_basic"

const (
	maxGap = 31
	maxN   = 1 << maxGap
)

var lg [maxN]int

func init() {
	incrBit := false
	for i := 1; i <= maxN; i++ {
		incrBit = 1<<lg[i-1] == i
		if incrBit {
			lg[i] = lg[i-1] + 1
		}
	}
}

type ST struct {
	skipGap  int
	initialN int
	stArr    [][]int

	isMaxST bool
}

func (st *ST) initialArr(intArr []int) {
	intArrL := len(intArr)
	mGap := lg[intArrL+1]
	st.stArr = make([][]int, intArrL+1, intArrL+1)
	for index := 0; index <= intArrL; index++ {
		st.stArr[index] = make([]int, mGap, mGap)
		st.stArr[index][0] = intArr[index]
	}
	for gap := 1; (1 << gap) <= intArrL; gap++ {
		for index := 1; index+(1<<gap)-1 <= intArrL; index++ {
			st.stArr[index][gap] = algo_basic.MaxValue(st.stArr[index][gap-1], st.stArr[index+1<<(gap-1)][gap-1])
		}
	}

}

func (st *ST) queryRange(startIndex, endIndex int) int {
	tGap := lg[endIndex-startIndex+1]
	return algo_basic.MaxValue(st.stArr[startIndex][tGap], st.stArr[endIndex-(1<<tGap)][tGap])
}

func (st *ST) ConstructST(arr []int) {
	st.initialArr(arr)
}
func (st *ST) QueryRange(l, r int) int {
	return st.QueryRange(l, r)
}
