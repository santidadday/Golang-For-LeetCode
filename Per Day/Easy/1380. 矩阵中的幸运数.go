package Easy

func luckyNumbers(matrix [][]int) []int {
	result := make([]int, 0)
	//行数为len(matrix)
	//列数为len(matrix[0])
	if matrix == nil {
		return nil
	}
	m := len(matrix)
	n := len(matrix[0])
	tag := 0
	for i := 0; i < m; i++ {
		count := matrix[i][0]
		//x记录count的列数，在找到行最小值时通过遍历列确定是否为列最大
		x := 0
		for j := 1; j < n; j++ {
			if matrix[i][j] < count {
				count = matrix[i][j]
				x = j
			}
		}
		for k := 0; k < m; k++ {
			if matrix[k][x] > count { //不符合列最大
				tag = -1
				break
			}
		}
		if tag == 0 {
			result = append(result, count)
		}
		tag = 0
	}
	return result
}
