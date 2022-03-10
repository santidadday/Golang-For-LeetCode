package Medium

import "fmt"

func complexNumberMultiply(num1 string, num2 string) string {
	//实部为实部相乘+虚部相乘
	//虚部为实部乘虚部

	len1 := len(num1)
	len2 := len(num2)

	//得到实部数值
	real1 := realGet(num1, len1)
	real2 := realGet(num2, len2)

	//得到虚部数值
	imaginary1 := imaginaryGet(num1, len1)
	imaginary2 := imaginaryGet(num2, len2)

	real := real1*real2 - imaginary1*imaginary2
	imaginary := real1*imaginary2 + real2*imaginary1
	result := fmt.Sprintf("%d", real) + "+" + fmt.Sprintf("%d", imaginary) + "i"

	return result
}

func realGet(num string, len int) int {
	real := 0
	if num[0] != '-' {
		for i := 0; num[i] != '+'; i++ {
			real = 10*real + int(num[i]) - 48
		}
	} else {
		for i := 1; num[i] != '+'; i++ {
			real = 10*real + int(num[i]) - 48
		}
		real = -real
	}
	return real
}

func imaginaryGet(num string, len int) int {
	count := 0
	imaginary := 0
	for num[count] != '+' {
		count++
	}

	if num[count+1] != '-' {
		for i := count + 1; num[i] != 'i'; i++ {
			imaginary = imaginary*10 + int(num[i]) - 48
		}
	} else {
		for i := count + 2; num[i] != 'i'; i++ {
			imaginary = imaginary*10 + int(num[i]) - 48
		}
		imaginary = -imaginary
	}
	return imaginary
}
