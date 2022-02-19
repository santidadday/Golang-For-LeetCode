package Easy

func findCenter(edges [][]int) int {
	//遍历edges[i]前两位即可得出中间节点出现次数为2
	result := 0

	count1 := edges[0][0]
	count2 := edges[0][1]

	if count1 == edges[1][0] || count1 == edges[1][1] {
		result = count1
	} else {
		result = count2
	}
	return result

}
