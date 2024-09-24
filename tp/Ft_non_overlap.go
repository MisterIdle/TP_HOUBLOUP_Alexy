package tp

import "sort"

//La fonction retourne le plus petit nombre d'intervalle à retirer pour que les intervalles
//restant ne se superpose pas.

func Ft_non_overlap(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	count := 1
	end := intervals[0][1]
	for _, interval := range intervals {
		start := interval[0]
		if start >= end {
			count++
			end = interval[1]
		}
	}
	return len(intervals) - count
}
