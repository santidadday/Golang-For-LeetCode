package Medium

func singleNonDuplicate(nums []int) int {
	//时间复杂度为log n只能使用折半查找方式
	//折半查找
	//mid为奇数时,若nums[mid]=nums[mid+1]，则X在mid右侧
	//mid为偶数时,若nums[mid]=nums[mid-1]，则X在mid右侧

	low := 0
	high := len(nums) - 1
	result := 0
	if high == 0 || nums[low] != nums[low+1] {
		result = nums[low]
		return result
	}
	if nums[high] != nums[high-1] {
		result = nums[high]
		return result
	}
	for low <= high {
		mid := (low + high) / 2
		if mid%2 == 0 {
			if nums[mid] == nums[mid+1] {
				low = mid + 1
			} else if nums[mid] == nums[mid-1] {
				high = mid - 1
			} else {
				result = nums[mid]
				break
			}
		} else {
			if nums[mid] == nums[mid-1] {
				low = mid + 1
			} else if nums[mid] == nums[mid+1] {
				high = mid - 1
			} else {
				result = nums[mid]
				break
			}
		}
	}
	return result
}
