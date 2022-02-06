package Easy

func twoSum(nums []int, target int) []int {
	//通过创建一个切片来接收遍历时返回的下标
	result := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result[0] = i
				result[1] = j
				return result
			}
		}
	}
	return nil
}
