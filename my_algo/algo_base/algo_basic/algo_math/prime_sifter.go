package algo_math

import "time"

const (
	defaultLength = 1e6 + 1
)

var primeArr = make([]int, defaultLength, defaultLength)

func SetPrimeArr(maxPrime int) {
	if maxPrime > defaultLength {
		primeArr = make([]int, maxPrime+1, maxPrime+1)
	}
	for base := 2; base <= maxPrime; base++ {
		for multiple := base << 1; multiple <= maxPrime; multiple += base {
			primeArr[multiple] = 1
		}
	}
}

func IsPrime(Number int) bool {
	if Number < len(primeArr) {
		return primeArr[Number] == 0 && Number != 1
	} else {
		panic("Number is beyond the Boundary! ")
	}
}

// CalculateOrderPrime
// 如果范围内没有该次序的素数，返回-1
func CalculateOrderPrime(start, end, TH int) int {
	if start < 1 && end > len(primeArr) {
		return -1
	}
	number, th := start, 0
	for number <= end {
		if IsPrime(number) {
			th++
		}
		if th == TH {
			return number
		}
		number++
	}
	return -1
}

// APrimeGame 一些人围成一圈，然后依次说出一个素数，假设上个人说出的素数是primeNum
// 那么下一个说出的素数必须在(primeNum,primeNum*2)之间，超时者或者说错者drink
func APrimeGame(personAmount, StartNum int, deadTime time.Time) {

}
