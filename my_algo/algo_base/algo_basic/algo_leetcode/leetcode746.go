package algo_leetcode

func Min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func minCostClimbingStairs(cost []int) int {
	l := len(cost)
	dp := make([]int, l+1)
	dp[0], dp[1] = cost[0], cost[1]
	for i := 2; i <= l; i++ {
		dp[i] = Min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[l]
}
