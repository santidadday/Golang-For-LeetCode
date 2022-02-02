package Easy

import "fmt"

func isPalindrome(x int) bool {
	//当传入数字为负数时，该数一定不是回文数
	if x < 0 {
		return false
	}
	//通过将传入数字转换为字符串形式对字符串进行遍历
	str := fmt.Sprintf("%d", x)
	for i := 0; i < len(str)/2; i++ {
		//注意：对字符串进行遍历是对字符串边界的限制，以防超出字符串边界
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}
	return true
}
