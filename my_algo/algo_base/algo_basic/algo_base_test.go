package algo_basic

import (
	"fmt"
	"github.com/Juminiy/my_go_lib/my_algo/algo_base/algo_basic/algo_math"
	"github.com/Juminiy/my_go_lib/my_algo/algo_base/algo_basic/algo_string"
	"strings"
	"testing"
)

func TestMySort_MyQuickSort(t *testing.T) {
	sort := &MySort{true}
	fmt.Println(sort.MyMergeSort([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, 0, 9))
}

func TestMyAbs(t *testing.T) {
	fmt.Println(algo_string.KMP("kksspsspsspss", "sp"))
}

func TestKMP(t *testing.T) {
	// match,index := algo_string.KMP("AAAAAABCD","ABCD")
	str1 := "AERSXFFRSXRSX"
	str1 = strings.Replace(str1, "RSX", "oppt", 2)
	algo_math.SetPrimeArr(1000)
	fmt.Println(algo_math.IsPrime(56))
	fmt.Println(algo_math.CalculateOrderPrime(1, 1000, 5))
}
