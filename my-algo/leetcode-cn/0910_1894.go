package leetcode_cn

import "github.com/Juminiy/my_go_lib/my-algo/algo_base/data_struct/simple"

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

//
func checkValidString(s string) bool {
	stack, count := &simple.MyStack{}, 0
	for _, ch := range s {
		if ch == '(' {
			stack.Push(ch)
		} else if ch == ')' {
			if !stack.IsEmpty() {
				if tCh, err := stack.Top(); err == nil {
					if tCh == '(' {
						stack.Pop()
					} else {
						if count > 0 {
							count--
						} else {
							return false
						}
					}
				} else {
					return false
				}
			} else {
				if count > 0 {
					count--
				} else {
					return false
				}
			}
		} else {
			count++
		}
	}
	for !stack.IsEmpty() {
		if tCh, err := stack.Top(); err == nil {
			if tCh == ')' {
				break
			} else {
				if count > 0 {
					count--
					stack.Pop()
				} else {
					break
				}
			}
		}
	}
	return stack.IsEmpty()
}
