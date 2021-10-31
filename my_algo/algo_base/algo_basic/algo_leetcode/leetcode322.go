package algo_leetcode

import "log"

func DebugPrint(x ...interface{}) {
	log.Println(x...)
}

// coinChange 方程思路是对的,没写全 168/188
func coinChange(coins []int, amount int) int {
	l := len(coins)
	dp := make([]int, amount+1)
	for i, _ := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for j := 0; j < l; j++ {
			if i >= coins[j] {
				dp[i] = Min(dp[i-coins[j]]+1, dp[i])
			}
		}

	}
	if dp[amount] > amount {
		return -1
	} else {
		return dp[amount]
	}
}
