package Medium

func countHighestScoreNodes(parents []int) (ans int) {
	maxScore, n, graph := int64(0), len(parents), map[int][]int{}
	for i := 1; i < n; i++ {
		graph[parents[i]] = append(graph[parents[i]], i)
	}

	var dfs func(node int) int
	dfs = func(node int) int {
		//三种情况:
		//1.若删除节点为根节点，则为两个数相乘
		//2.若删除节点为叶子节点，则只有一个树
		//3.若删除节点即存在父节点也存在孩子节点，则根据孩子节点个数。
		//若仅有一个孩子节点，则为两数相乘；若有两个孩子节点则为三数相乘

		//对上面三种情况进行简化，无论出现何种情况都会出现三个数。left right (n - left - right - 1)
		//不足1的使用以代替
		left, right := 0, 0
		children := graph[node]
		if len(children) > 0 {
			left = dfs(children[0])
			if len(children) > 1 {
				right = dfs(children[1])
			}
		}
		if score := max(left, 1) * max(right, 1) * max(n-1-left-right, 1); score > maxScore {
			maxScore, ans = score, 1
		} else if score == maxScore {
			ans++
		}
		return left + right + 1
	}
	dfs(0)
	return ans
}

func max(a, b int) int64 {
	if a > b {
		return int64(a)
	}
	return int64(b)
}
