package Medium

//使用顺序查找方式,时间复杂度为O(n)
func searchRange1(nums []int, target int) []int {
	result := make([]int, 2)
	result[0] = -1
	result[1] = -1
	Maxlen := len(nums)
	if Maxlen == 0 {
		return result
	}
	if nums[0] == target {
		result[0] = 0
	}
	if nums[Maxlen-1] == target {
		result[1] = Maxlen - 1
	}
	for i := 1; i < Maxlen; i++ {
		if nums[i-1] != target && nums[i] == target {
			result[0] = i
		}
		if nums[i] != target && nums[i-1] == target {
			result[1] = i - 1
		}
	}
	return result
}

//使用二分查找方式,时间复杂度为O(log n)
func searchRange(nums []int, target int) []int {
	result := make([]int, 2)
	result[0] = -1
	result[1] = -1
	if len(nums) == 0 {
		return result
	}
	low := 0
	high := len(nums) - 1
	mid := (low + high) / 2
	for low <= high {
		if nums[mid] > target {
			high = mid - 1
			mid = (low + high) / 2
		} else if nums[mid] < target {
			low = mid + 1
			mid = (low + high) / 2
		} else if nums[mid] == target {
			result[0] = mid
			result[1] = mid
			l1 := low
			h1 := mid - 1
			m1 := (l1 + h1) / 2

			l2 := mid + 1
			h2 := high
			m2 := (l2 + h2) / 2

			for l1 <= h1 {
				if nums[m1] < target {
					l1 = m1 + 1
					m1 = (l1 + h1) / 2
				} else {
					if m1 < result[0] {
						result[0] = m1
						h1 = m1 - 1
						m1 = (l1 + h1) / 2
					}
				}
			}

			for l2 <= h2 {
				if nums[m2] > target {
					h2 = m2 - 1
					m2 = (l2 + h2) / 2
				} else {
					if m2 > result[1] {
						result[1] = m2
						l2 = m2 + 1
						m2 = (l2 + h2) / 2
					}
				}
			}
			break
		}
	}
	return result
}
