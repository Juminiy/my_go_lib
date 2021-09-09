package algo_basic

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

type MySort struct {
	isAsc bool
}

var tArr []int = make([]int, maxArrLen, maxArrLen)

func (mySort *MySort) myComp(compValue, compedValue int) bool {
	if mySort.isAsc {
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
		for mySort.myComp(arr[j], pivot) {
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
		tArr[l1] = arr[ll]
		l1++
	}
	for ; l2 <= r2; ll++ {
		tArr[l2] = arr[ll]
		l2++
	}
	ll = l
	for ll = l; ll <= r; ll++ {
		arr[ll] = tArr[ll]
	}
}
