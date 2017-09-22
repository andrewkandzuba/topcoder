package datatypes

type ListNode struct {
	v int
	n *ListNode
}

func (ls ListNode) value() int {
	return ls.v
}

func (ls ListNode) next() *ListNode {
	return ls.n
}