package Medium

func generateParenthesis(n int) []string {
	//我们可以将所有情况都列出来,然后对其是否可行进行判断
	result := make([]string, 0)

	dfs(0, 0, n, "", &result)
	return result
}

func dfs(left, right, n int, s string, result *[]string) {
	if left == right && left == n {
		*result = append(*result, s)
		return
	}
	if left > n || right > n || right > left {
		return
	}
	dfs(left+1, right, n, s+"(", result)
	dfs(left, right+1, n, s+")", result)
}
