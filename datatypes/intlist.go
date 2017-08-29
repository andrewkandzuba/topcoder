package datatypes

type IntListNode struct {
	v int
	n ListNode
}

func (listNode IntListNode) value() int {
	return listNode.v
}

func (listNode IntListNode) next() ListNode {
	return listNode.n
}
