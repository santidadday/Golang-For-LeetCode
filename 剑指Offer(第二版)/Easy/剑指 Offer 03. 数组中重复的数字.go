package Easy

func findRepeatNumber(nums []int) int {
	//首先为暴力枚举法，通过两个for循环可以轻松找到重复数字
	//算法优化，通过空间换时间方式。先建立一个map，通过循环遍历并将出现次数不断添加至map
	//数组(切片)遍历完成之后，对map进行遍历若出现其中之大于1的情况可返回其值

	m := make(map[int]int, 0)

	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}

	for k, v := range m {
		if v > 1 {
			return k
		}
	}

	return -1
}
