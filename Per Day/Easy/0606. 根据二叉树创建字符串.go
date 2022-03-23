package Easy

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func tree2str(root *TreeNode) string {
	//先序遍历首先加入"("在遍历结束该节点之后加入")"
	result := ""
	if root != nil {
		result = result + fmt.Sprintf("%d", root.Val)
		if root.Left != nil {
			sort(root.Left, &result)
			if root.Right != nil {
				sort(root.Right, &result)
			}
		}
		if root.Left == nil && root.Right != nil {
			result = result + "()"
			sort(root.Right, &result)
		}
	}
	return result

}

func sort(root *TreeNode, result *string) {
	*result = *result + "("
	if root != nil {
		*result = *result + fmt.Sprintf("%d", root.Val)
	}
	if root.Left != nil {
		sort(root.Left, result)
		if root.Right != nil {
			sort(root.Right, result)
		}
	}
	if root.Left == nil && root.Right != nil {
		*result = *result + "()"
		sort(root.Right, result)
	}
	*result = *result + ")"
	return
}
