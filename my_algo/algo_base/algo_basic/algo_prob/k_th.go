package algo_prob

import "errors"

// algo comes from cs course
// also in https://leetcode-cn.com/problems/kth-largest-element-in-an-array/

func selectRange(nums []int) ([]int, int, []int, error) {
	if len(nums) <= 0 {
		return nil, 0, nil, errors.New("nums is nil! ")
	}
	pi, lo, hi := nums[0], make([]int, 0), make([]int, 0)
	for i := 1; i < len(nums); i++ {
		if nums[i] > pi {
			hi = append(hi, nums[i])
		} else {
			lo = append(lo, nums[i])
		}
	}
	return lo, pi, hi, nil
}

func partition(nums []int, l, r int) int {
	pivot, i, j := nums[0], l+1, r
	for i < j {
		for j > l && pivot > nums[j] {
			j--
		}
		for i < r && pivot < nums[i] {
			i++
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[l], nums[j] = nums[j], nums[l]
	return i
}

// pass but not fast & mem vs!

func FindKthLargest(nums []int, k int) int {
	l, r := 0, len(nums)-1
	for {
		i := partition(nums, l, r)
		if i == k-1 {
			return nums[i-1]
		} else if i > k-1 {
			l = i + 1
		} else {
			r = i - 1
		}
	}
}
