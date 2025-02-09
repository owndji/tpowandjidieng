package tp

import (
	"math"
)

func Ft_coin(coins []int, amount int) int {
	dp := make([]int, amount+1)

	for i := range dp {
		dp[i] = math.MaxInt32
	}

	dp[0] = 0

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i-coin >= 0 {
				dp[i] = int(math.Min(float64(dp[i]), float64(dp[i-coin]+1)))
			}
		}
	}

	if dp[amount] == math.MaxInt32 {
		return -1
	}

	return dp[amount]
}
