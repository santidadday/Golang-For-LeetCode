package Medium

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	//使用双指针同时遍历，两个间出发的差距为n，当前指针指向最后一个节点时则删除后指针对应节点
	if head.Next == nil {
		return nil
	}

	var front *ListNode = head
	var last *ListNode = head
	//假头指针
	newhead := &ListNode{0, head}
	var p *ListNode = newhead
	count := 1

	for count < n {
		front = front.Next
		count++
	}

	for front.Next != nil {
		front = front.Next
		last = last.Next
		p = p.Next
	}

	if last.Next != nil {
		p.Next = last.Next
	} else {
		p.Next = nil
	}

	return newhead.Next
}
