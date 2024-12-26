package n0005

import (
	"strings"
)

func longestPalindrome(s string) string {
	t := strings.Join(strings.Split("@"+s+"$", ""), "#")
	p := manacher(t)

	maxPalindromeLength := 0
	bestCenter := -1

	for i, v := range p {
		if v > maxPalindromeLength {
			maxPalindromeLength = v
			bestCenter = i
		}
	}

	l := (bestCenter - maxPalindromeLength) / 2
	r := (bestCenter + maxPalindromeLength) / 2

	return s[l:r]
}

func manacher(t string) []int {
	p := make([]int, len(t))

	center := 0

	for i := 1; i < len(t)-1; i++ {
		rightBoundary := center + p[center]
		mirrorIndex := center - (i - center)
		if rightBoundary > i {
			p[i] = min(rightBoundary-i, p[mirrorIndex])
		}
		for t[i+1+p[i]] == t[i-1-p[i]] {
			p[i] += 1
		}
		if i+p[i] > rightBoundary {
			center = 1
		}
	}

	return p
}
