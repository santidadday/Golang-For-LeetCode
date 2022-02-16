package Medium

func lengthOfLongestSubstring(s string) int {
	result := 0
	n := len(s)
	//从每个字母开始遍历
	//若某个字符和前方已加入字符冲突则停止统计count与result对比
	//优化措施，当发生冲突时，直接从冲突字后开始计算，并且count为两个count下标相减
	if n <= 1 {
		return n
	}
	//i为起始点下标，j为当前下标，k遍历从i到j
	for i := 0; i < n; i++ {
		count := 1
		tag := 0
		for j := i + 1; j < n; j++ {
			for k := i; k < j; k++ {
				if s[k] == s[j] {
					tag = 1
					i = k
					break
				}
			}
			if tag == 1 {
				break
			} else {
				count++
			}
		}
		if count > result {
			result = count
		}
	}
	return result
}
