package Medium

import "math"

//BFS  2.11
func letterCombinations(digits string) []string {
	//使用广度优先求解问题
	//定义一个管道，从digits读取数据
	//1.如果管道为空，则直接加入管道
	//2.如果管道不空，则从管道中弹出一个字符，把当前字母对应的字符与弹出字符相组合后加入管道
	//当digits遍历完成之后，遍历管道并加入最终的切片中

	Maxlen := len(digits)
	caps := int(math.Pow(4, float64(Maxlen)))
	ch := make(chan string, caps)

	result := make([]string, 0)
	count := 0
	if Maxlen > 0 {
		switch digits[0] {
		case '2':
			ch <- "a"
			ch <- "b"
			ch <- "c"
			count = 3
		case '3':
			ch <- "d"
			ch <- "e"
			ch <- "f"
			count = 3
		case '4':
			ch <- "g"
			ch <- "h"
			ch <- "i"
			count = 3
		case '5':
			ch <- "j"
			ch <- "k"
			ch <- "l"
			count = 3
		case '6':
			ch <- "m"
			ch <- "n"
			ch <- "o"
			count = 3
		case '7':
			ch <- "p"
			ch <- "q"
			ch <- "r"
			ch <- "s"
			count = 4
		case '8':
			ch <- "t"
			ch <- "u"
			ch <- "v"
			count = 3
		case '9':
			ch <- "w"
			ch <- "x"
			ch <- "y"
			ch <- "z"
			count = 4
		}
	}

	c := count

	for i := 1; i < len(digits); i++ {
		if digits[i] == '2' || digits[i] == '3' || digits[i] == '4' || digits[i] == '5' || digits[i] == '6' || digits[i] == '8' {
			count = count * 3
		} else {
			count = count * 4
		}

		for j := 0; j < c; j++ {
			res := <-ch
			if digits[i] == '2' {
				ch <- (res + "a")
				ch <- (res + "b")
				ch <- (res + "c")
			} else if digits[i] == '3' {
				ch <- (res + "d")
				ch <- (res + "e")
				ch <- (res + "f")
			} else if digits[i] == '4' {
				ch <- (res + "g")
				ch <- (res + "h")
				ch <- (res + "i")
			} else if digits[i] == '5' {
				ch <- (res + "j")
				ch <- (res + "k")
				ch <- (res + "l")
			} else if digits[i] == '6' {
				ch <- (res + "m")
				ch <- (res + "n")
				ch <- (res + "o")
			} else if digits[i] == '7' {
				ch <- (res + "p")
				ch <- (res + "q")
				ch <- (res + "r")
				ch <- (res + "s")
			} else if digits[i] == '8' {
				ch <- (res + "t")
				ch <- (res + "u")
				ch <- (res + "v")
			} else {
				ch <- (res + "w")
				ch <- (res + "x")
				ch <- (res + "y")
				ch <- (res + "z")
			}
		}

		c = count
	}

	close(ch)

	for v := range ch {
		result = append(result, v)
	}
	return result
}
