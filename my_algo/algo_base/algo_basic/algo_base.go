package algo_basic

import (
	"math"
	"strconv"
)

// 算法经过 https://luogu.com.cn/problem/P1177 测试

const (
	maxBit        = 31
	maxArrLen     = 1e6
	notFoundIndex = -1
)

func MaxValue(compValue, compedValue int) int {
	if compValue > compedValue {
		return compValue
	} else {
		return compedValue
	}
}
func MinValue(compValue, compedValue int) int {
	if compValue < compedValue {
		return compValue
	} else {
		return compedValue
	}
}

func StrToInt(str string) int {
	if strInt, err := strconv.Atoi(str); err == nil {
		return strInt
	}
	return math.MaxInt64
}
func IntToStr(intValue int) string {
	return strconv.Itoa(intValue)
}
func IntToChar(intNum int) rune {
	if intNum <= 9 {
		return rune(intNum + 48)
	} else {
		panic("intValue must <= 9! ")
	}
}
func CharToInt(char rune) int {
	if char <= 57 {
		return int(char - 48)
	} else {
		panic("char must range 49~57! ")
	}
}
func MyAbs(value int) int {
	if value >= 0 {
		return value
	} else {
		return -1 * value
	}
}

type MySort struct {
	IsAsc bool
}

var tArr []int = make([]int, maxArrLen, maxArrLen)
var tStorage [][]int = make([][]int, maxBit, maxArrLen)

func (mySort *MySort) myComp(compValue, compedValue int) bool {
	if mySort.IsAsc {
		return compValue < compedValue
	} else {
		return compValue > compedValue
	}
}
func (mySort *MySort) MyQuickSort(arr []int, l, r int) {
	if l > r {
		return
	}
	i, j, pivot := l, r, arr[(r-l)/2+l]
	for i <= j {
		for mySort.myComp(pivot, arr[j]) {
			j--
		}
		for mySort.myComp(arr[i], pivot) {
			i++
		}
		if i <= j {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}
	}
	if i < r {
		mySort.MyQuickSort(arr, i, r)
	}
	if l < j {
		mySort.MyQuickSort(arr, l, j)
	}
}

var tLevels = 0

func (mySort *MySort) div(arr []int, l, r int) {
	if l >= r {
		return
	}
	mid := (r-l)/2 + l
	mySort.div(arr, l, mid)
	mySort.div(arr, mid+1, r)
	mySort.merge(arr, l, r)
	tStorage[tLevels] = tArr[:r+1]
	tLevels++
}
func (mySort *MySort) merge(arr []int, l, r int) {
	mid := (r-l)/2 + l
	l1, r1, l2, r2, ll := l, mid, mid+1, r, l
	for ; l1 <= r1 && l2 <= r2; ll++ {
		if mySort.myComp(arr[l1], arr[l2]) {
			tArr[ll] = arr[l1]
			l1++
		} else {
			tArr[ll] = arr[l2]
			l2++
		}
	}
	for ; l1 <= r1; ll++ {
		tArr[ll] = arr[l1]
		l1++
	}
	for ; l2 <= r2; ll++ {
		tArr[ll] = arr[l2]
		l2++
	}
	ll = l
	for ll = l; ll <= r; ll++ {
		arr[ll] = tArr[ll]
	}
}
func (mySort *MySort) MyMergeSort(arr []int, l, r int) [][]int {
	mySort.div(arr, l, r)
	return tStorage[:tLevels]
}

// BinarySearch must be sorted Arr else Meaningless
// but the algo is not accurate when the same element occurs over twice times.
func BinarySearch(arr []int, l, r, value int) int {
	mid := 0
	for l <= r {
		mid = (r-l)/2 + l
		if arr[mid] == value {
			return mid
		} else if arr[mid] > value {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return notFoundIndex
}

// LowerBound from l find the first arr[i] <= value
// example. {10,10,20,20,20,20,30,30},0,7,20 return index is 2
// algo complex is O(log(r-l))
func LowerBound(arr []int, l, r, value int) int {
	dist, step, cur := r-l+1, 0, 0
	for dist > 0 {
		cur, step, cur = l, dist/2, step+cur
		if arr[cur] < value {
			l, dist = cur+1, dist-step-1
		} else {
			dist = step
		}
	}
	return l
}

// UpperBound from r find the first arr[i] > value
// example. {10,10,20,20,20,20,30,30},0,7,20 return index is 6
// algo complex is O(log(r-l))
func UpperBound(arr []int, l, r, value int) int {
	dist, step, cur := r-l+1, 0, 0
	for dist > 0 {
		cur, step, cur = l, dist/2, step+cur
		if !(value < arr[cur]) {
			l, dist = cur+1, dist-step-1
		} else {
			dist = step
		}
	}
	return l
}
