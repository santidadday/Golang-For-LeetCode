package Easy

func isOneBitCharacter(bits []int) bool {
	if bits[len(bits)-1] == 1 {
		return false
	}
	//除最后一位外,剩余字符必须有10、11、0组成
	//组成情况,其可由任意个0相连，偶数个1相连。若由奇数个1相连其后面必须至少有一个0且该0不是最后一位
	//所以 我们的算法统计bits中有多少个1相连

	count := 0
	for i := 0; i < len(bits)-1; i++ {
		if bits[i] == 1 {
			count++
		} else if count%2 == 0 {
			count = 0
		} else if count%2 != 0 && bits[i] == 0 {
			count = 0
		}
	}
	if count%2 != 0 {
		return false
	}
	return true
}
