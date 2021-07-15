package main

func coinChange(coins []int, amount int) int {
	// 需要创建 amount+1 大小的空间
	// dp 保存当 amount = i时，dp[i]表示需要的硬币最少
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		dp[i] = amount + 1
	}
	for i := 1; i <= amount; i++ {
		// F(n)=min(F(n−1),F(n−4),F(n−5))+1
		// 根据动态转移方程，依次填充dp[i]的值
		// 后面的dp[i]依赖于前面的[i]数值
		for j := 0; j < len(coins); j++ {
			// F(n-1) => dp[i-coins[0]]
			// F(n-4) => dp[i-coins[1]]
			// F(n-5) => dp[i-coins[2]]
			if coins[j] <= i {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
