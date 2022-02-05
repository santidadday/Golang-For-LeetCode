package Middle

import (
	"fmt"
	"math"
	"strconv"
)

func reverse(x int) int {
	//环境无法存储64位整数，所有无法使用移位运算进行
	//通过字符串进行转换
	if x > 0 {
		//首先将整数转为字符串，其次对字符串进行运算，最后将字符串转换为整数
		str := fmt.Sprintf("%d", x)
		Maxlen := len(str)
		s := []rune(str)
		for i := 0; i < Maxlen/2; i++ {
			//交换
			s[i], s[Maxlen-i-1] = s[Maxlen-i-1], s[i]
		}
		r, _ := strconv.ParseInt(string(s), 10, 64)
		result := int(r)
		if float64(result) > math.Pow(2, 31)-1 {
			return 0
		} else {
			return result
		}

	} else if x == 0 {
		return 0
	} else {
		str := fmt.Sprintf("%d", x)
		Maxlen := len(str)
		s := []rune(str)
		for i := 0; i < (Maxlen-1)/2; i++ {
			s[i+1], s[Maxlen-i-1] = s[Maxlen-i-1], s[i+1]
		}
		r, _ := strconv.ParseInt(string(s), 10, 64)
		result := int(r)
		if float64(result) < math.Pow(-2, 31) {
			return 0
		} else {
			return result
		}
	}
}
