package tp

// Ft_min_window retourne la plus petite séquence dans "s" qui contient tous les caractères de "t".
// Si deux séquences ont la même taille, la séquence apparaissant en premier dans "s" est retournée.

func Ft_min_window(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}

	// Fréquence des caractères dans t
	m := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		m[t[i]]++
	}

	left := 0
	right := 0
	match := 0

	start := 0
	minLen := len(s) + 1

	windowCount := make(map[byte]int)

	for right < len(s) {
		charRight := s[right]
		if _, ok := m[charRight]; ok {
			windowCount[charRight]++
			if windowCount[charRight] == m[charRight] {
				match++
			}
		}

		for match == len(m) {
			if right-left+1 < minLen {
				minLen = right - left + 1
				start = left
			}

			charLeft := s[left]
			if _, ok := m[charLeft]; ok {
				windowCount[charLeft]--
				if windowCount[charLeft] < m[charLeft] {
					match--
				}
			}
			left++
		}

		right++
	}

	if minLen == len(s)+1 {
		return ""
	}

	return s[start : start+minLen]
}
