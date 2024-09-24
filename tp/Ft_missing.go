package tp

// Ft_missing retourne l'unique nombre manquant dans l'intervalle [0,n]

func Ft_missing(nums []int) int {
	n := len(nums)
	expectedSum := (n * (n + 1)) / 2
	actualSum := 0

	for _, num := range nums {
		actualSum += num
	}

	return expectedSum - actualSum
}
