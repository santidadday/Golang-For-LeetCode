package Easy

import "strings"

func divideString(s string, k int, fill byte) []string {
	if len(s)%k == 0 {
		//无需扩容增加fill
		var result []string
		if k != 1 {
			result = make([]string, (len(s)+1)/k)
		} else {
			result = make([]string, len(s)/k)
		}
		slice := s[:]
		var index int = 0
		for i := 0; i < len(s)+1-k; i = i + k {
			result[index] = slice[i : i+k]
			index++
		}
		return result
	} else {
		//需要扩容
		slice := strings.Split(s, "")
		var result []string = make([]string, (len(s)/k)+1)
		var index int = 0
		add := k*(len(s)/k+1) - len(s)
		for i := 0; i < add; i++ {
			slice = append(slice, string(fill))
		}
		st := strings.Join(slice, "")
		for i := 0; i < len(s)+add; i = i + k {
			result[index] = st[i : i+k]
			index++
		}
		return result
	}
}
