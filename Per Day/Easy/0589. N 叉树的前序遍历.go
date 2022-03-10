package Easy

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	//使用递归算法
	//对子节点进行遍历
	result := make([]int, 0)
	dfs(root, &result)

	return result
}

func dfs(node *Node, r *[]int) {
	if node == nil {
		return
	}
	*r = append(*r, node.Val)
	for _, ch := range node.Children {
		dfs(ch, r)
	}
}
