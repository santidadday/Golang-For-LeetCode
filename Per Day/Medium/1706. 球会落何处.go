package Medium

func findBall(grid [][]int) []int {
	n := len(grid[0])

	result := make([]int, n)
	for i := range result {
		count := i
		for _, row := range grid {
			dir := row[count]
			count += dir                                      // 移动球
			if count < 0 || count == n || row[count] != dir { // 到达侧边或 V 形
				count = -1
				break
			}
		}
		result[i] = count // col >= 0 为成功到达底部
	}

	return result
}
