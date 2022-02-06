package Easy

//通过map将七种字符所代表的的整数值存入其中
var symbolValues = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

func romanToInt(s string) int {
	var result int = 0
	//对s字符串进行遍历
	for i, _ := range s {
		value := symbolValues[s[i]]
		//对于最后一位元素一定是加入结果中
		if i < len(s)-1 && value < symbolValues[s[i+1]] {
			result = result - value
		} else {
			result = result + value
		}
	}
	return result
}
