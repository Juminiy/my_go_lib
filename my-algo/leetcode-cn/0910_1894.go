package leetcode_cn

func chalkReplacer(chalk []int, k int) int {
	n, i, seqSum := len(chalk), 0, 0
	for ; i < n; i++ {
		seqSum += chalk[i]
	}
	k %= seqSum
	for i = 0; k >= chalk[i]; i++ {
		k -= chalk[i]
	}
	return i
}
