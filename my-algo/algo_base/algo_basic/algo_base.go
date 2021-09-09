package algo_basic

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

func MyMergeSort(arr []int) {

}
