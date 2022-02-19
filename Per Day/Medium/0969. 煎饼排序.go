package Medium

func pancakeSort(arr []int) []int {
	result := make([]int, 0)
	//从后向前排序，翻转后需要完成将当前存在最大的数置于尾部
	//每一个数都是互不相同的，且最小值为1最大值为arr.length
	//输出结果不唯一
	Maxlen := len(arr)
	for Maxlen > 1 { //每次将最大数置于尾部后，就可以放弃该数，对剩余数字进行访问
		for i := 0; i < Maxlen; i++ {
			if arr[i] == Maxlen && i != 0 { //数字唯一，且每次的最大数等于该次长度
				result = append(result, i+1)
				//变换数组
				temp(arr, i+1)
				result = append(result, Maxlen)
				temp(arr, Maxlen)
				Maxlen--
				break
			} else if arr[i] == Maxlen && i == 0 {
				result = append(result, Maxlen)
				temp(arr, Maxlen)
				Maxlen--
				break
			}
		}
	}
	return result
}

func temp(arr []int, count int) {
	for i := 0; i < count/2; i++ {
		arr[i], arr[count-i-1] = arr[count-i-1], arr[i]
	}
}
