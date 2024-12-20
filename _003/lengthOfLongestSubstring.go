package _003

func lengthOfLongestSubstring(s string) int {
	ans := 0

	count := make([]int, 128)

	l := 0
	for r, c := range s {
		count[c] += 1
		for count[c] > 1 {
			count[s[l]] -= 1
			l += 1
		}
		ans = max(ans, r-l+1)
	}

	return ans
}
