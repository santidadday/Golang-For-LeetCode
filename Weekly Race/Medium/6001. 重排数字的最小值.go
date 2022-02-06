package Medium

import (
	"math"
	"sort"
)

func smallestNumber(num int64) int64 {
	n1 := int(num)
	s := make([]int, 0)
	if n1 == 0 {
		return 0
	}
	if n1 > 0 {
		for {
			if n1 == 0 {
				break
			} else {
				n := n1 % 10
				s = append(s, n)
			}
			n1 = n1 / 10
		}
		sort.Ints(s)
		if s[0] == 0 {
			for i := 0; i < len(s); i++ {
				if s[i] != 0 {
					s[i], s[0] = s[0], s[i]
					break
				}
			}
		}
		result := 0
		for i := len(s) - 1; i >= 0; i-- {
			result = result + s[i]*int(math.Pow(10, float64(len(s)-1-i)))
		}
		return int64(result)

	} else {
		//取绝对值
		n1 = int(math.Abs(float64(num)))

		for {
			if n1 == 0 {
				break
			} else {
				n := n1 % 10
				s = append(s, n)
			}
			n1 = n1 / 10
		}
		sort.Ints(s)
		Maxlen := len(s)
		for i := 0; i < Maxlen/2; i++ {
			s[i], s[Maxlen-i-1] = s[Maxlen-i-1], s[i]
		}

		result := 0
		for i := len(s) - 1; i >= 0; i-- {
			result = result + s[i]*int(math.Pow(10, float64(len(s)-1-i)))
		}
		return int64(-result)
	}
}
