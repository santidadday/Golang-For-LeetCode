package Medium

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeNodes(head *ListNode) *ListNode {
	var p *ListNode = head
	var q *ListNode = head
	count := 0

	for p.Next != nil {
		if p.Next.Val != 0 {
			count = count + p.Next.Val
			p = p.Next
		} else {
			q.Next.Val = count
			q = q.Next
			q.Next = p.Next
			p = p.Next
			q = p
			count = 0
		}
	}

	p = head.Next
	for p != nil {
		p.Next = p.Next.Next
		p = p.Next
	}

	return head.Next
}
