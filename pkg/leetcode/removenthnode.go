package leetcode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	remove(head, n, 0)
	return head
}

func remove(head *ListNode, n int, pos int) int {
	if head != nil {
		removeAt := remove(head.Next, n, pos+1)
		if removeAt == pos {
			if head.Next != nil && head.Next.Next != nil {
				head.Next = head.Next.Next
			} else {
				head.Next = nil
			}
		}
		return removeAt
	}
	return pos - n
}