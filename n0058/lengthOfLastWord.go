package n0058

func lengthOfLastWord(s string) int {
	i := len(s) - 1
	for i >= 0 && s[i] == ' ' {
		i--
	}
	lastIndex := i
	for i >= 0 && s[i] != ' ' {
		i--
	}
	return lastIndex - i
}
