package Easy

import "fmt"

func countEven(num int) int {
	sum := 0
	for i := 1; i <= num; i++ {
		str := fmt.Sprintf("%d", i)
		count := 0
		for j := 0; j < len(str); j++ {
			count = int(str[j]) + count
		}
		if count%2 == 0 {
			sum++
		}
	}
	return sum
}
