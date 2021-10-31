package algo_leetcode

// leetcode136
func singleNumber(nums []int) int {
	base := 0
	for i := 0; i < len(nums); i++ {
		base ^= nums[i]
	}
	return base
}

// leetcode260
func singleNumberX(nums []int) []int {
	base := 0
	for i := 0; i < len(nums); i++ {
		base ^= nums[i]
	}
	tBase, x1, x2 := base&-base, 0, 0
	for _, x := range nums {
		if tBase&x > 0 {
			x1 ^= x
		} else {
			x2 ^= x
		}
	}
	return []int{x1, x2}
}
