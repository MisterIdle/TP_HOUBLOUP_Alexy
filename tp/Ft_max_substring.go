package tp

//La valeur de retour est la longueur de la plus grande sous chaine de s possédant des
//caractères non répétés au sein de s.

func Ft_max_substring(s string) int {
	m := make(map[byte]int)
	left := 0
	res := 0
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; ok {
			left = max(left, m[s[i]]+1)
		}
		m[s[i]] = i
		res = max(res, i-left+1)
	}
	return res
}
