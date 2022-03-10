package Easy

import "unicode"

func reverseOnlyLetters(s string) string {
	//仅有英文字母翻转，双指针调换

	slice := []byte(s)
	left, right := 0, len(s)-1
	for {
		// 判断左边是否扫描到字母
		for left < right && !unicode.IsLetter(rune(s[left])) {
			left++
		}
		// 判断右边是否扫描到字母
		for right > left && !unicode.IsLetter(rune(s[right])) {
			right--
		}
		if left >= right {
			break
		}
		slice[left], slice[right] = slice[right], slice[left]
		left++
		right--
	}
	return string(slice)
}
