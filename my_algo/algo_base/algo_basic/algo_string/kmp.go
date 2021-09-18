package algo_string

const (
	StringNotMatch = -1
)

// BruteForceMatch index from 0~len(S)-len(P)
func BruteForceMatch(S, P string) (bool, int) {
	if len(P) <= len(S) {
		for i := 0; i <= len(S)-len(P); i++ {
			if S[i:i+len(P)] == P {
				return true, i
			}
		}
	}
	return false, 0
}

func NextArr(P string) []int {
	nextArr, index, cur := make([]int, len(P)+1), 1, 0
	for index < len(P) {
		if P[index] == P[cur] {
			cur++
			nextArr[index] = cur
			index++
		} else if cur != 0 {
			cur = nextArr[cur-1]
		} else {
			index++
			nextArr[index] = cur
		}
	}
	return nextArr
}

// KMP string matching is the basic task of a computer
// Knuth-Morris-Pratt
func kmp(MainStr, PatternStr string) []int {
	tar, pos, nextArr, resArr := 0, 0, NextArr(PatternStr), make([]int, 0)
	for tar < len(MainStr) {
		if MainStr[tar] == PatternStr[pos] {
			tar++
			pos++
		} else if pos != 0 {
			pos = nextArr[pos-1]
		} else {
			tar++
		}
		if pos == len(PatternStr) {
			resArr = append(resArr, tar-pos)
			pos = nextArr[pos-1]
		}
	}
	return resArr
}
func KMP(S, P string) (bool, int) {
	resArr := kmp(S, P)
	if len(resArr) > 0 {
		return true, resArr[0]
	} else {
		return false, StringNotMatch
	}
}
func KMPGlobal(S, P string) (bool, []int) {
	resArr := kmp(S, P)
	if len(resArr) > 0 {
		return true, resArr
	} else {
		return false, nil
	}
}
