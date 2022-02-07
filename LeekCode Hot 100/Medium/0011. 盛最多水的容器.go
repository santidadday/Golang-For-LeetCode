package Medium

//使用双指针(头、尾)算法代替双层for循环
//双层for循环在提交时会报超时
func maxArea(height []int) int {
	var left int = 0
	var right int = len(height) - 1
	result := 0
	for left != right {
		count := (right - left) * min(left, right, height)
		if result < count {
			result = count
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return result
}

//找出i，j下表中最小的数
func min(a, b int, height []int) int {
	if height[a] >= height[b] {
		return height[b]
	} else {
		return height[a]
	}
}
