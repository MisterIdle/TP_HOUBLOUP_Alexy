package main

import (
	"fmt"
	"tp/tp"
)

func main() {
	fmt.Println("=========================================")
	fmt.Println("FT_COIN")
	fmt.Println(tp.Ft_coin([]int{1, 2, 5}, 11)) // Resultat : 3
	fmt.Println(tp.Ft_coin([]int{2}, 3))        // Resultat : -1
	fmt.Println(tp.Ft_coin([]int{1}, 0))        // Resultat : 0

	fmt.Println("\nFT_MISSING")
	fmt.Println(tp.Ft_missing([]int{3, 1, 2}))                   // Resultat : 0
	fmt.Println(tp.Ft_missing([]int{0, 1}))                      // Resultat : 2
	fmt.Println(tp.Ft_missing([]int{9, 6, 4, 2, 3, 5, 7, 0, 1})) // Resultat : 8

	fmt.Println("\nFT_NON_OVERLAP")
	fmt.Println(tp.Ft_non_overlap([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}})) // Resultat : 1
	fmt.Println(tp.Ft_non_overlap([][]int{{1, 2}, {2, 3}}))                 // Resultat : 0
	fmt.Println(tp.Ft_non_overlap([][]int{{1, 2}, {1, 2}, {1, 2}}))         // Resultat : 2

	fmt.Println("\nFT_PROFIT")
	fmt.Println(tp.Ft_profit([]int{7, 1, 5, 3, 6, 4})) // Resultat : 5
	// si on achète au jour 1, nous payons 1,
	// et si nous le vendons au 4eme jour, nous gagnons 6, le bénéfice est 6-1
	fmt.Println(tp.Ft_profit([]int{7, 6, 4, 3, 1})) // Resultat : 0

	fmt.Println("\nFT_MAX_SUBSTRING")
	fmt.Println(tp.Ft_max_substring("abcabcbb")) // Resultat : 3
	// "abc" est la plus grande sous chaine composé de caractère diffèrent
	fmt.Println(tp.Ft_max_substring("bbbbb")) // Resultat : 1
	// "b" est la plus grande sous chaine

	fmt.Println("\nFT_MIN_WINDOW")
	fmt.Println(tp.Ft_min_window("ADOBECODEBANC", "ABC")) // Resultat : "BANC"
	fmt.Println(tp.Ft_min_window("a", "aa"))              // Resultat : ""
	fmt.Println("=========================================")
}
