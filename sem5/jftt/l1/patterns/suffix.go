package patterns

func Suffix(pattern, text string) int {
	maxLen := 0
	minLen := len(pattern)
	if len(text) < minLen {
		minLen = len(text)
	}
	for i := 1; i <= minLen; i++ {
		prefix := pattern[:i]
		suffix := text[len(text)-i:]
		if prefix == suffix {
			maxLen = i
		}
	}
	return maxLen
}
