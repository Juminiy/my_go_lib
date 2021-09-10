package algo_basic

// 算法经过 https://luogu.com.cn/problem/P1177 测试，保证正确性

const (
	maxBit    = 30
	maxArrLen = 1e6
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

func (mySort *MySort) MyMergeSort(arr []int, l, r int) {
	if l >= r {
		return
	}
	mid := (r-l)/2 + l
	mySort.MyMergeSort(arr, l, mid)
	mySort.MyMergeSort(arr, mid+1, r)

	l1, r1, l2, r2, ll := l, mid, mid+1, r, l
	for ; l1 <= r1 && l2 <= r2; ll++ {
		if arr[l1] < arr[l2] {
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

func BinarySearch(arr []int, l, r, value int) int {
	mid := 0
	for l < r {
		mid = (r-l)/2 + l
		if arr[mid] == value {
			return mid
		} else if arr[mid] > value {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

// TODO
func LowerBound(arr []int, l, r, value int) int {
	return BinarySearch(arr, l, r, value)
}

// TODO
func UpperBound(arr []int, l, r, value int) int {
	return BinarySearch(arr, l, r, value)
}
