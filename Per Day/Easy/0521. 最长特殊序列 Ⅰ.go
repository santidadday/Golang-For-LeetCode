package Easy

func findLUSlength(a string, b string) int {
	if a != b {
		return max(len(a), len(b))
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
