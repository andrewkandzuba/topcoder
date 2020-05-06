package leetcode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	_, head = remove(head, head,nil, n, 1)
	return head
}

func remove(head *ListNode, curr *ListNode, prev *ListNode, n int, len int) (int, *ListNode) {
	if curr != nil {
		removeAt, _ := remove(head, curr.Next, curr, n, len+1)
		if removeAt == len {
			if prev == nil {
				head = curr.Next
			} else {
				prev.Next = curr.Next
			}
		}
		return removeAt, head
	}
	return len - n, nil
}
