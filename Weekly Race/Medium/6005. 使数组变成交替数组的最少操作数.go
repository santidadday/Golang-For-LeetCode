package Medium

func minimumOperations(nums []int) int {
	result := 0
	//记录M1第一多和第二多
	x1 := 0
	xn1 := 0
	x2 := 0
	//记录M2第一多和第二多
	y1 := 0
	yn1 := 0
	y2 := 0
	m1 := make(map[int]int, 0)
	m2 := make(map[int]int, 0)

	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			m1[nums[i]]++
		} else {
			m2[nums[i]]++
		}
	}

	for k, v := range m1 {
		if v >= x1 {
			x2 = x1
			x1 = v

			xn1 = k
		} else if v >= x2 && v < x1 {
			x2 = v

		}
	}
	for k, v := range m2 {
		if v >= y1 {
			y2 = y1
			y1 = v

			yn1 = k
		} else if v >= y2 && v < y1 {
			y2 = v

		}
	}
	if xn1 != yn1 {
		result = len(nums) - x1 - y1
	} else if xn1 == yn1 && x2 > y2 {
		result = len(nums) - y1 - x2
	} else {
		result = len(nums) - x1 - y2
	}

	return result
}
