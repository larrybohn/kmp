package kmp

import "fmt"

func FindMatches(text string, pattern string) []int {
	fmt.Printf("text: %v\npattern: %v\n", text, pattern)

	lps := make([]int, len(pattern))
	res := make([]int, 0)

	lps[0] = 0
	for i, j := 0, 1; j < len(pattern); {
		if pattern[i] == pattern[j] {
			lps[j] = i + 1
			i++
			j++
		} else if i == 0 {
			lps[j] = 0
			j++
		} else {
			i = lps[i-1]
		}
	}

	fmt.Printf("lps: %v\n", lps)

	for i, j := 0, 0; i < len(text); {
		if text[i] == pattern[j] {
			i++
			j++
			if j == len(pattern) {
				res = append(res, i-j)
				j = lps[j-1]
			}
		} else if j > 0 {
			j = lps[j-1]
		} else {
			i++
		}
	}

	return res
}
