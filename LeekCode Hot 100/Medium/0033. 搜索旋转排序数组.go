package Medium

//使用顺序查找方式,时间复杂度为O(log n)
func search1(nums []int, target int) int {
	Maxlen := len(nums)
	if target > nums[Maxlen-1] {
		//在正面查找
		for i := 0; i < Maxlen; i++ {
			if nums[i] > target || (i > 0 && nums[i] < nums[i-1]) {
				return -1
			}
			if nums[i] == target {
				return i
			}
		}
	} else {
		for i := Maxlen - 1; i > -1; i-- {
			if nums[i] < target || (i < Maxlen-1 && nums[i+1] < nums[i]) {
				return -1
			}
			if nums[i] == target {
				return i
			}
		}
	}
	return -1
}

//使用二分查找方式,时间复杂度为O(log n)
func search2(nums []int, target int) int {
	//使用类似于折半搜索方式进行搜索,可将时间复杂度降低至O(log n)
	low := 0
	high := len(nums) - 1
	mid := (low + high) / 2
	if nums[high] < target {
		//在0~k-1之间寻找
		for low <= high {
			if nums[mid] < nums[0] {
				high = mid - 1
				mid = (low + high) / 2
			} else {
				if nums[mid] > target {
					high = mid - 1
					mid = (low + high) / 2
				} else if nums[mid] < target {
					low = mid + 1
					mid = (low + high) / 2
				} else if nums[mid] == target {
					return mid
				}
			}
		}
	} else if nums[low] > target {
		//在k~n-1之间寻找
		for low <= high {
			if nums[mid] > nums[len(nums)-1] {
				low = mid + 1
				mid = (low + high) / 2
			} else {
				if nums[mid] > target {
					high = mid - 1
					mid = (low + high) / 2
				} else if nums[mid] < target {
					low = mid + 1
					mid = (low + high) / 2
				} else if nums[mid] == target {
					return mid
				}
			}
		}
	} else if nums[low] == target || nums[high] == target {
		if nums[low] == target {
			return low
		} else {
			return high
		}
	} else if nums[0] < nums[len(nums)-1] {
		//序列有序，直接使用二分查找即可
		for low <= high {
			if nums[mid] > target {
				high = mid - 1
				mid = (low + high) / 2
			} else if nums[mid] < target {
				low = mid + 1
				mid = (low + high) / 2
			} else if nums[mid] == target {
				return mid
			}
		}
	}
	return -1
}
