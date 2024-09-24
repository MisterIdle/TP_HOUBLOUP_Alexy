package tp

//La fonction retourne le plus petit nombre de pièces nécessaire pour atteindre la valeur
//amount selon les pièces précisées dans coins.

func Ft_coin(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if coin <= i {
				dp[i] = min(dp[i], 1+dp[i-coin])
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}

	return dp[amount]
}
