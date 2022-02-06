package Medium

func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	//申请两个字符用于记录起始位置及最大长度
	var maxlen int = 1
	var index int = 0
	//第一层循环为确定起始位置
	//第二层循环为确定最大回文子串的最大长度
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if j-i+1 > maxlen && JudgePalindrome(s, i, j) {
				maxlen = j - i + 1
				index = i
			}
		}
	}
	return s[index : index+maxlen]
}

func JudgePalindrome(s string, left int, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
