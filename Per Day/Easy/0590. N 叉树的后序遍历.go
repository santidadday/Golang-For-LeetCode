package Easy

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	result := make([]int, 0)
	//后序遍历：如果有孩子节点，则继续递归遍历孩子节点；如果没有孩子节点或其孩子节点均被遍历完成，则将该节点加入最终返回切片
	dfs(root, &result)
	return result
}

func dfs(root *Node, result *[]int) {
	if root == nil {
		return
	}
	for i := 0; i < len(root.Children); i++ {
		dfs(root.Children[i], result)
	}
	*result = append(*result, root.Val)
}
