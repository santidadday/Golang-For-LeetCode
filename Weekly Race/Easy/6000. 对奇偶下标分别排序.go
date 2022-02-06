package Easy

func sortEvenOdd(nums []int) []int {
	MaxLen := len(nums)
	for i := 0; i < (MaxLen-1)/2; i++ {
		j := 0
		for {
			if j < MaxLen && j+2 < MaxLen {
				if nums[j] > nums[j+2] {
					nums[j], nums[j+2] = nums[j+2], nums[j]
				}
			} else {
				break
			}
			j = j + 2
		}
	}

	for i := 0; i < (MaxLen-1)/2; i++ {
		j := 1
		for {
			if j < MaxLen && j+2 < MaxLen {
				if nums[j] < nums[j+2] {
					nums[j], nums[j+2] = nums[j+2], nums[j]
				}
			} else {
				break
			}
			j = j + 2
		}
	}
	return nums
}
