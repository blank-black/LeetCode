package problems

func CoinChange(coins []int, amount int) int {
	max := amount + 1
	dp := make([]int, max)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = max
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				dp[i] = min(dp[i], dp[i - coins[j]] + 1)
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

